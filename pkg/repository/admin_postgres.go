package repository

import (
	"github.com/aalmat/e-commerce/models"
	"github.com/jinzhu/gorm"
	"time"
)

type AdminPostgres struct {
	db *gorm.DB
}

func (a *AdminPostgres) UpdateProduct(productId uint, update models.ProductUpdate) error {
	err := update.Validate()
	if err != nil {
		return err
	}

	var product models.Product
	if err := a.db.Where("id = ?", productId).First(&product).Error; err != nil {
		return err
	}

	if update.Title != "" {
		product.Title = update.Title
	}
	if update.Description != "" {
		product.Description = update.Description
	}
	if update.Photo != "" {
		product.Photo = update.Photo
	}

	if err := a.db.Save(&product).Error; err != nil {
		return err
	}

	return nil
}

func (a *AdminPostgres) GetProducts() ([]models.Product, error) {
	var products []models.Product
	if err := a.db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (a *AdminPostgres) CreateProduct(product models.Product) (uint, error) {
	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()
	product.Rating = 0
	product.Quantity = 0
	if err := a.db.Select("title", "description", "photo", "created_at", "updated_at", "rating", "quantity").Create(&product).Error; err != nil {
		return 0, err
	}

	return product.ID, nil
}

func (a *AdminPostgres) DeleteProduct(productId uint) error {
	err := a.db.Where("id = ?", productId).Delete(models.Product{}).Error
	return err
}

func NewAdminPostgres(db *gorm.DB) *AdminPostgres {
	return &AdminPostgres{db}
}
