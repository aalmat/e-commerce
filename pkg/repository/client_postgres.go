package repository

import (
	"fmt"
	"github.com/aalmat/e-commerce/models"
	"github.com/jinzhu/gorm"
	"strings"
)

type ClientPostgres struct {
	db *gorm.DB
}

func (c *ClientPostgres) FilterByPrice(minPrice, maxPrice int) ([]models.WareHouse, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ClientPostgres) FilterByRating(minRate, maxRate int) ([]models.WareHouse, error) {
	//TODO implement me
	panic("implement me")
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

func (c *ClientPostgres) AddToCart(userId uint, productId uint, quantity uint) (uint, error) {
	tx := c.db.Begin()
	item := models.Cart{UserID: userId, ProductID: productId, Quantity: quantity}
	if err := tx.Select("user_id", "product_id", "quantity").Create(&item).Error; err != nil {
		tx.Rollback()
		return 0, err
	}
	if _, err := c.ChangeProductQuantity(item.UserID, item.ProductID, item.Quantity); err != nil {
		tx.Rollback()
		return 0, err
	}
	return item.ID, nil
}

func (c *ClientPostgres) ShowCartProducts(userId uint) ([]models.WareHouse, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ClientPostgres) DeleteFromCart(userId, productId uint) error {
	//TODO implement me
	panic("implement me")
}

func (c *ClientPostgres) SearchByName(keyword string) ([]models.Product, error) {
	words := strings.Split(keyword, " ")
	query := ""
	for _, v := range words {
		query += "%"
		query += v
		query += "% "
	}

	var products []models.Product
	if err := c.db.Where(fmt.Sprintf("producs.name LIKE %s", query)).Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil

}
