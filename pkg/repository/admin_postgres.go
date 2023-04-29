package repository

import (
	"github.com/aalmat/e-commerce/models"
	"github.com/jinzhu/gorm"
)

type AdminPostgres struct {
	db *gorm.DB
}

func (a AdminPostgres) CreateProduct(product models.Product) (uint, error) {
	//TODO implement me
	panic("implement me")
}

func (a AdminPostgres) DeleteProduct(productId uint) error {
	//TODO implement me
	panic("implement me")
}

func NewAdminPostgres(db *gorm.DB) *AdminPostgres {
	return &AdminPostgres{db}
}
