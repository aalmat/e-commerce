package repository

import (
	"github.com/aalmat/e-commerce/models"
	"github.com/jinzhu/gorm"
)

type ProductPostgres struct {
	db *gorm.DB
}

func (p ProductPostgres) GetAll() ([]models.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductPostgres) GetById(userId uint, productId uint) (models.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductPostgres) SearchByName(keyword string) ([]models.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductPostgres) FilterByPrice(minPrice, maxPrice int) ([]models.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductPostgres) FilterByRating(minRate, maxRate int) ([]models.Product, error) {
	//TODO implement me
	panic("implement me")
}

func NewProductPostgres(db *gorm.DB) *ProductPostgres {
	return &ProductPostgres{db}
}
