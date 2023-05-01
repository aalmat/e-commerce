package repository

import (
	"errors"
	"github.com/aalmat/e-commerce/models"
	"github.com/jinzhu/gorm"
)

type ClientPostgres struct {
	db *gorm.DB
}

func (c *ClientPostgres) ChangeProductQuantity(userid uint, productId uint, quantity uint) (uint, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ClientPostgres) FilterByPrice(minPrice, maxPrice int) ([]models.WareHouse, error) {
	//TODO implement me
	panic("implement me")
}

func NewClientPostgres(db *gorm.DB) *ClientPostgres {
	return &ClientPostgres{db}
}

func (c *ClientPostgres) AddToCart(userId uint, productId uint, quantity uint) (uint, error) {
	tx := c.db.Begin()
	item := models.Cart{}
	if err := tx.Where("user_id = ? AND product_id = ?", userId, productId).First(&item).Error; err != nil {
		if !gorm.IsRecordNotFoundError(err) {
			tx.Rollback()
			return 0, err
		}
		// товар жок болса новый запись
		item = models.Cart{UserID: userId, ProductID: productId, Quantity: quantity}
		if err := tx.Create(&item).Error; err != nil {
			tx.Rollback()
			return 0, err
		}
	} else {
		// товар бар болса то quantity кобееди
		item.Quantity += quantity
		if err := tx.Save(&item).Error; err != nil {
			tx.Rollback()
			return 0, err
		}
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return 0, err
	}
	return item.ID, nil
}

func (c *ClientPostgres) ShowCartProducts(userId uint) ([]models.WareHouse, error) {
	var warehouse []models.WareHouse
	err := c.db.
		Table("cart").
		Select("warehouse.*, cart.quantity as cart_quantity").
		Joins("inner join warehouse on warehouse.product_id = cart.product_id").
		Where("cart.user_id = ?", userId).
		Scan(&warehouse).
		Error
	if err != nil {
		return nil, err
	}
	return warehouse, nil
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
func (c *ClientPostgres) SearchByName(keyword string) ([]models.WareHouse, error) {
	var warehouse []models.WareHouse
	err := c.db.Table("warehouse").
		Joins("JOIN product ON product.id = warehouse.product_id").
		Where("product.name LIKE ?", "%"+keyword+"%").
		Find(&warehouse).
		Error
	if err != nil {
		return nil, err
	}
	return warehouse, nil
}

func (c *ClientPostgres) FilterByRating(minRate, maxRate int) ([]models.WareHouse, error) {
	tx := c.db.Begin()

	var items []models.WareHouse
	if err := tx.Where("rating BETWEEN ? AND ?", minRate, maxRate).Find(&items).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return items, nil
}

func (c *ClientPostgres) RateProduct(userId, productId uint, rate uint) (uint, error) {
	tx := c.db.Begin()

	// Проверяем, существует ли запись рейтинга для данного пользователя и продукта
	var existingRating models.Rating
	if err := tx.Where("user_id = ? AND product_id = ?", userId, productId).First(&existingRating).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			tx.Rollback()
			return 0, err
		}
	}

	// Если запись рейтинга существует, обновляем ее, бомаса создаем новую запись
	if existingRating.ID != 0 {
		existingRating.Rate = rate
		if err := tx.Save(&existingRating).Error; err != nil {
			tx.Rollback()
			return 0, err
		}
	} else {
		newRating := models.Rating{
			Rate:      rate,
			UserId:    userId,
			ProductId: productId,
		}
		if err := tx.Create(&newRating).Error; err != nil {
			tx.Rollback()
			return 0, err
		}
	}

	// Рассчитываем новый средний рейтинг товара и обновляем его
	var product models.Product
	if err := tx.First(&product, productId).Error; err != nil {
		tx.Rollback()
		return 0, err
	}

	var totalRatings uint
	var ratingsCount uint
	if err := tx.Model(&models.Rating{}).Where("product_id = ?", productId).Count(&ratingsCount).Error; err != nil {
		tx.Rollback()
		return 0, err
	}

	if ratingsCount > 0 {
		if err := tx.Model(&models.Rating{}).Where("product_id = ?", productId).Select("SUM(rate)").Row().Scan(&totalRatings); err != nil {
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
func (c *ClientPostgres) WriteComment(userId, productId uint, commentText string) (uint, error) {
	tx := c.db.Begin()

	var user models.User
	if err := tx.First(&user, userId).Error; err != nil {
		tx.Rollback()
		return 0, err
	}

	var product models.Product
	if err := tx.First(&product, productId).Error; err != nil {
		tx.Rollback()
		return 0, err
	}

	comment := models.Commentary{
		UserId:    userId,
		ProductId: productId,
		Text:      commentText,
	}
	if err := tx.Create(&comment).Error; err != nil {
		tx.Rollback()
		return 0, err
	}

	if err := tx.Commit().Error; err != nil {
		return 0, err
	}

	return comment.ID, nil
}
