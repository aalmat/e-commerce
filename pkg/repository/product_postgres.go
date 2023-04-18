package repository

import (
	"github.com/aalmat/e-commerce/models"
	"github.com/jinzhu/gorm"
)

type ProductPostgres struct {
	db *gorm.DB
}

func NewProductPostgres(db *gorm.DB) *ProductPostgres {
	return &ProductPostgres{db: db}
}

func (p *ProductPostgres) GetAll() ([]models.Product, error) {
	var products []models.Product
	if err := p.db.Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}
func (p *ProductPostgres) CreateProduct(userId uint, product models.Product) (uint, error) { // product id, error
	
	return 0, nil
}
func (p *ProductPostgres) GetAllSellerProduct(userId uint) ([]models.Product, error) {
	return nil, nil
}
func (p *ProductPostgres) GetById(userId uint, productId uint) (models.Product, error) {
	return models.Product{}, nil
}
func (p *ProductPostgres) SearchByName(keyword string) ([]models.Product, error) {
	return nil, nil
}
func (p *ProductPostgres) FilterByPrice(minPrice, maxPrice int) ([]models.Product, error) {
	return nil, nil
}
func (p *ProductPostgres) FilterByRating(minRate, maxRate int) ([]models.Product, error) {
	return nil, nil
}
