package repository

import (
	"github.com/aalmat/e-commerce/models"
	"github.com/jinzhu/gorm"
	"strings"
)

type ProductPostgres struct {
	db *gorm.DB
}

func (p *ProductPostgres) GetAll() ([]models.Product, error) {
	var products []models.Product
	if err := p.db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (p *ProductPostgres) GetById(productId uint) (models.Product, error) {
	var product models.Product
	if err := p.db.Where("id = ?", productId).First(&product).Error; err != nil {
		return models.Product{}, err
	}
	return product, nil
}

func (p *ProductPostgres) SearchByName(search models.Search) ([]models.Product, error) {
	words := strings.Split(search.Keyword, " ")
	query := ""
	for _, v := range words {
		if len(query) != 0 {
			query += " "
		}
		query += "%"
		query += v
		query += "%"
	}

	var products []models.Product
	if err := p.db.Where("title like ?", query).Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil

}

func NewProductPostgres(db *gorm.DB) *ProductPostgres {
	return &ProductPostgres{db}
}
