package repository

import (
	"github.com/aalmat/e-commerce/models"
	"github.com/jinzhu/gorm"
)

type AdminPostgres struct {
	db *gorm.DB
}

func (a *AdminPostgres) UpdateProduct(productId uint, update models.ProductUpdate) error {
	//TODO implement me
	panic("implement me")
	return nil
}

func (a *AdminPostgres) GetProducts() ([]models.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (a *AdminPostgres) CreateProduct(product models.Product) (uint, error) {
	//TODO implement me
	panic("implement me")
}

func (a *AdminPostgres) DeleteProduct(productId uint) error {
	//TODO implement me
	panic("implement me")
}

func NewAdminPostgres(db *gorm.DB) *AdminPostgres {
	return &AdminPostgres{db}
}
