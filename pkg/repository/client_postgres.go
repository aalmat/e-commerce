package repository

import (
	"github.com/aalmat/e-commerce/models"
	"github.com/jinzhu/gorm"
)

type ClientPostgres struct {
	db *gorm.DB
}

func (c *ClientPostgres) ChangeProductQuantity(userid uint, productId uint, quantity uint) (uint, error) {
	return 0, nil
}

func NewClientPostgres(db *gorm.DB) *ClientPostgres {
	return &ClientPostgres{db}
}

func (c *ClientPostgres) AddToCart(userId uint, productId uint, quantity uint) (uint, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ClientPostgres) ShowCartProducts(userId uint) ([]models.WareHouse, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ClientPostgres) DeleteFromCart(userId, productId uint) error {
	//TODO implement me
	panic("implement me")
}

func (c *ClientPostgres) SearchByName(keyword string) ([]models.WareHouse, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ClientPostgres) FilterByPrice(minPrice, maxPrice int) ([]models.WareHouse, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ClientPostgres) FilterByRating(minRate, maxRate int) ([]models.WareHouse, error) {
	//TODO implement me
	panic("implement me")
}
