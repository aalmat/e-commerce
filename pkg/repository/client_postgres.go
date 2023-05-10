package repository

import (
	"errors"
	"fmt"
	"github.com/aalmat/e-commerce/models"
	"github.com/jinzhu/gorm"
	"time"
)

type ClientPostgres struct {
	db *gorm.DB
}

const (
	deliveryInterval = 5
)

func (c *ClientPostgres) PurchaseAll(userId uint) error {
	tx := c.db.Begin()

	var whs []models.Cart
	if err := tx.Where("user_id = ?", userId).Find(&whs).Error; err != nil {
		return err
	}

	//fmt.Println("whs ", whs)

	for i := range whs {
		if err := c.PurchaseById(userId, whs[i].ID); err != nil {
			tx.Rollback()
			return err
		}
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}
	//if err := tx.Select("user_id", "product_id", "quantity")
	return nil
}

func (c *ClientPostgres) PurchaseById(userId uint, cartId uint) error {
	tx := c.db.Begin()

	var cart models.Cart
	if err := tx.Where("user_id=? and id=?", userId, cartId).First(&cart).Error; err != nil {
		return errors.New(fmt.Sprintf("error selecting cart with id %d", cart.ID))
	}

	Order := c.CartToOrder(cart)
	if err := c.db.Table("ware_houses").Select("price").Where("id = ? and user_id=?", cart.WareHouseID, cart.UserID).Scan(&Order.Price).Error; err != nil {
		return err
	}
	if err := tx.Select("user_id", "ware_house_id", "quantity", "delivery_date", "status", "created_at", "updated_at").Create(&Order).Error; err != nil {
		tx.Rollback()
		return errors.New(fmt.Sprintf("error adding product with is %d", Order.ID))
	}

	if err := c.ChangeWhQuantity(cart.WareHouseID, cart.Quantity); err != nil {
		tx.Rollback()
		return err
	}

	var product models.WareHouse
	if err := tx.Where("id=?", cart.WareHouseID).First(&product).Error; err != nil {
		return errors.New(fmt.Sprintf("error selecting cart with id %d", cart.ID))
	}

	if err := c.ChangePrQuantity(product.ProductId, cart.Quantity); err != nil {
		tx.Rollback()
		return err
	}

	if err := c.DeleteFromCart(userId, cartId); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil

}

func (c *ClientPostgres) ChangeWhQuantity(productId, subs uint) error {
	var wh models.WareHouse
	if err := c.db.Where("id=?", productId).First(&wh).Error; err != nil {
		return err
	}

	if wh.Quantity < subs {
		return errors.New("products in warehouse is less than you want")
	}

	wh.Quantity = wh.Quantity - subs
	if err := c.db.Save(&wh).Error; err != nil {
		return err
	}
	return nil
}
func (c *ClientPostgres) ChangePrQuantity(productId, subs uint) error {
	var wh models.Product
	if err := c.db.Where("id=?", productId).First(&wh).Error; err != nil {
		return err
	}

	if wh.Quantity < subs {
		return errors.New("products quantity is less than you want")
	}

	wh.Quantity = wh.Quantity - subs
	if err := c.db.Save(&wh).Error; err != nil {
		return err
	}
	return nil
}

func (c *ClientPostgres) CartToOrder(wh models.Cart) models.Order {
	var Order models.Order

	hours := time.Now().Sub(wh.CreatedAt).Hours()
	days := int(hours / 24)
	days = days % deliveryInterval
	h := time.Duration(days * 24)

	Order.DeliveryDate = time.Now().Add(time.Hour * h)
	Order.CreatedAt = time.Now()
	Order.Quantity = wh.Quantity
	Order.UpdatedAt = wh.UpdatedAt
	Order.WareHouseID = wh.WareHouseID
	Order.UserID = wh.UserID
	Order.Status = false

	return Order
}

func (c *ClientPostgres) RateProduct(rate models.Rating) (uint, error) {
	tx := c.db.Begin()

	// Проверяем, существует ли запись рейтинга для данного пользователя и продукта
	var existingRating models.Rating
	if err := tx.Where("user_id = ? AND product_id = ?", rate.UserId, rate.ProductId).First(&existingRating).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			tx.Rollback()
			return 0, err
		}
	}

	// Если запись рейтинга существует, обновляем ее, бомаса создаем новую запись
	if existingRating.ID != 0 {
		existingRating.Rate = rate.Rate
		if err := tx.Save(&existingRating).Error; err != nil {
			tx.Rollback()
			return 0, err
		}
	} else {

		if err := tx.Create(&rate).Error; err != nil {
			tx.Rollback()
			return 0, err
		}
	}

	// Рассчитываем новый средний рейтинг товара и обновляем его
	var product models.Product
	if err := tx.First(&product, rate.ProductId).Error; err != nil {
		tx.Rollback()
		return 0, err
	}

	var totalRatings uint
	var ratingsCount uint
	if err := tx.Model(&models.Rating{}).Where("product_id = ?", rate.ProductId).Select("COUNT(product_id)").Row().Scan(&ratingsCount); err != nil {
		tx.Rollback()
		return 0, err
	}
	if ratingsCount > 0 {
		if err := tx.Model(&models.Rating{}).Where("product_id = ?", rate.ProductId).Select("SUM(rate)").Row().Scan(&totalRatings); err != nil {
			tx.Rollback()
			return 0, err
		}
		product.Rating = totalRatings / ratingsCount
	} else {
		product.Rating = 0
	}
	if err := tx.Save(&product).Error; err != nil {
		tx.Rollback()
		return 0, err
	}
	if err := tx.Commit().Error; err != nil {
		return 0, err
	}
	return product.Rating, nil
}

func (c *ClientPostgres) WriteComment(comment models.Commentary) (uint, error) {
	tx := c.db.Begin()

	var product models.Product
	if err := tx.First(&product, comment.ProductId).Error; err != nil {
		tx.Rollback()
		return 0, err
	}

	comment.CreatedAt = time.Now()
	comment.UpdatedAt = time.Now()

	if err := tx.Select("user_id", "product_id", "text", "created_at", "updated_at").Create(&comment).Error; err != nil {
		tx.Rollback()
		return 0, err
	}

	if err := tx.Commit().Error; err != nil {
		return 0, err
	}

	return comment.ID, nil
}

func (c *ClientPostgres) ChangeProductQuantity(userid uint, cartId uint, quantity uint) (uint, error) {
	var cartItem models.Cart
	cartItem.UserID = userid
	cartItem.ID = cartId
	if err := c.db.First(&cartItem).Error; err != nil {
		return 0, nil
	}
	cartItem.Quantity = quantity
	err := c.db.Save(&cartItem).Error
	if err != nil {
		return 0, nil
	}
	return cartItem.ID, nil
}

func NewClientPostgres(db *gorm.DB) *ClientPostgres {
	return &ClientPostgres{db}
}

func (c *ClientPostgres) AddToCart(userId uint, whId uint, quantity uint) (uint, error) {
	tx := c.db.Begin()

	var wh models.WareHouse
	if err := tx.Where("id = ?", whId).Find(&wh).Error; err != nil {
		tx.Rollback()
		return 0, err
	}

	if wh.Quantity < quantity {
		tx.Rollback()
		return 0, errors.New("ware house has less product quantity")
	}

	var item models.Cart
	err := tx.Where("user_id=? and ware_house_id=?", userId, wh.ID).First(&item).Error
	if !gorm.IsRecordNotFoundError(err) && err != nil {
		tx.Rollback()
		return 0, err
	}

	if gorm.IsRecordNotFoundError(err) {

		item = models.Cart{UserID: userId, WareHouseID: wh.ID, Quantity: quantity}
		item.UpdatedAt = time.Now()
		item.CreatedAt = time.Now()
		if err := tx.Select("user_id", "ware_house_id", "quantity", "created_at", "updated_at").Create(&item).Error; err != nil {
			tx.Rollback()
			return 0, err
		}
		if err := tx.Commit().Error; err != nil {
			tx.Rollback()
			return 0, err
		}
		return item.ID, nil
	} else {
		return 0, errors.New("you already added this product to cart")
	}

}

func (c *ClientPostgres) ShowCartProducts(userId uint) ([]models.CartInfo, error) {
	var whs []models.CartInfo

	if err := c.db.Table("ware_houses").Select("distinct on (ware_house_id) carts.id as cart_id, ware_houses.product_id, ware_houses.user_id, ware_houses.price, ware_houses.created_at, ware_houses.id, carts.quantity").Joins("inner join carts on ware_houses.id = carts.ware_house_id").Where("carts.user_id=? and carts.deleted_at is null", userId).Scan(&whs).Error; err != nil {
		return nil, err
	}

	return whs, nil
}

func (c *ClientPostgres) DeleteFromCart(userId, cartId uint) error {
	tx := c.db.Begin()
	cart := models.Cart{}

	if err := tx.Where("user_id = ? AND id = ?", userId, cartId).Delete(&cart).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (c *ClientPostgres) ShowOrders(userId uint) ([]models.Order, error) {
	var orders []models.Order
	if err := c.db.Where("user_id = ?", userId).Order("created_at desc").Find(&orders).Error; err != nil {
		return nil, err
	}

	return orders, nil
}
