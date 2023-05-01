package repository

import (
	"errors"
	"fmt"
	"github.com/aalmat/e-commerce/models"
	"github.com/jinzhu/gorm"
	"time"
)

var defaultDeliveryDate time.Time

type ClientPostgres struct {
	db *gorm.DB
}

func (c *ClientPostgres) PurchaseAll(userId uint) error {
	tx := c.db.Begin()

	var whs []models.Cart
	if err := tx.Where("user_id = ?", userId).Find(&whs).Error; err != nil {
		return err
	}

	fmt.Println("whs ", whs)

	for i := range whs {
		Order := c.CartToOrder(whs[i])

		if err := tx.Select("user_id", "product_id", "quantity", "delivery_date", "status", "created_at", "updated_at").Create(&Order).Error; err != nil {
			tx.Rollback()
			return errors.New(fmt.Sprintf("error adding product with is %d %s", Order.ID, err.Error()))
		}

		if err := c.ChangeWhQuantity(whs[i].ProductID, whs[i].Quantity); err != nil {
			tx.Rollback()
			return err
		}
		if err := c.ChangePrQuantity(whs[i].ProductID, whs[i].Quantity); err != nil {
			tx.Rollback()
			return err
		}

		if err := c.DeleteFromCart(userId, whs[i].ProductID); err != nil {
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

func (c *ClientPostgres) PurchaseById(userId uint, productIds []uint) error {
	tx := c.db.Begin()

	for i := range productIds {
		var wh models.Cart
		if err := tx.Where("user_id=? and product_id=?", userId, productIds[i]).First(&wh).Error; err != nil {
			return errors.New(fmt.Sprintf("error selecting cart with id %d", wh.ID))
		}

		Order := c.CartToOrder(wh)

		if err := tx.Select("user_id", "product_id", "quantity", "delivery_date", "status", "created_at", "updated_at").Create(Order).Error; err != nil {
			tx.Rollback()
			return errors.New(fmt.Sprintf("error adding product with is %d", Order.ID))
		}

		if err := c.ChangeWhQuantity(productIds[i], wh.Quantity); err != nil {
			tx.Rollback()
			return err
		}
		if err := c.ChangePrQuantity(productIds[i], wh.Quantity); err != nil {
			tx.Rollback()
			return err
		}

		if err := c.DeleteFromCart(userId, productIds[i]); err != nil {
			tx.Rollback()
			return err
		}
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil

}

func (c *ClientPostgres) ChangeWhQuantity(productId, subs uint) error {
	var wh models.WareHouse
	if err := c.db.Where("product_id=?", productId).First(&wh).Error; err != nil {
		return err
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
	days = days % 5
	h := time.Duration(days * 24)

	Order.DeliveryDate = time.Now().Add(time.Hour * h)
	Order.CreatedAt = time.Now()
	Order.Quantity = wh.Quantity
	Order.UpdatedAt = wh.UpdatedAt
	Order.ProductID = wh.ProductID
	Order.UserID = wh.UserID
	Order.Status = false

	return Order
}

func (c *ClientPostgres) RateProduct(userId, productId uint, rate uint) (uint, error) {
	return 0, nil
}

func (c *ClientPostgres) WriteComment(userId, productId uint, commentText string) (uint, error) {
	return 0, nil
}

func (c *ClientPostgres) ChangeProductQuantity(userid uint, productId uint, quantity uint) (uint, error) {
	var cartItem models.Cart
	cartItem.UserID = userid
	cartItem.ProductID = productId
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
	err := tx.Where("user_id=? and product_id=?", userId, wh.ProductId).First(&item).Error
	if !gorm.IsRecordNotFoundError(err) && err != nil {
		tx.Rollback()
		return 0, err
	}

	if gorm.IsRecordNotFoundError(err) {

		item = models.Cart{UserID: userId, ProductID: wh.ProductId, Quantity: quantity}
		item.UpdatedAt = time.Now()
		item.CreatedAt = time.Now()
		if err := tx.Select("user_id", "product_id", "quantity", "created_at", "updated_at").Create(&item).Error; err != nil {
			tx.Rollback()
			return 0, err
		}
		if err := tx.Commit().Error; err != nil {
			tx.Rollback()
			return 0, err
		}
		return item.ID, nil
	} else {
		id, err1 := c.ChangeProductQuantity(item.UserID, item.ProductID, item.Quantity)
		if err1 != nil {
			tx.Rollback()
			return 0, err1
		}
		if err := tx.Commit().Error; err != nil {
			tx.Rollback()
			return 0, err
		}
		return id, nil
	}

}

func (c *ClientPostgres) ShowCartProducts(userId uint) ([]models.WareHouse, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ClientPostgres) DeleteFromCart(userId, productId uint) error {
	tx := c.db.Begin()

	if err := tx.Where("user_id = ? AND product_id = ?", userId, productId).Delete(&models.Cart{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
