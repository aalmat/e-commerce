package repository

import (
	"github.com/aalmat/e-commerce/models"
	"github.com/jinzhu/gorm"
)

type ClientPostgres struct {
	db *gorm.DB
}

func NewClientPostgres(db *gorm.DB) *ClientPostgres {
	return &ClientPostgres{db}
}

func (c ClientPostgres) AddToCart(productId uint, quantity uint) (uint, error) {
	//TODO implement me
	panic("implement me")
}

func (c ClientPostgres) ShowCartProducts() ([]models.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (c ClientPostgres) DeleteFromCart(productId uint, quantity uint) (uint, error) {
	//TODO implement me
	panic("implement me")
}
