package repository

import (
	"github.com/aalmat/e-commerce/models"
	"github.com/jinzhu/gorm"
)

type Authorization interface {
	CreateUser(user models.User) (uint, error)
	GetUser(email, password string) (models.User, error)
}

type Product interface {
	GetAll() ([]models.Product, error)
	CreateProduct(userId uint, product models.Product) (uint, error) // product id, error
	GetAllSellerProduct(userId uint) ([]models.Product, error)
	GetById(userId uint, productId uint) (models.Product, error)
	SearchByName(keyword string) ([]models.Product, error)
	FilterByPrice(minPrice, maxPrice int) ([]models.Product, error)
	FilterByRating(minRate, maxRate int) ([]models.Product, error)
}

type Repository struct {
	Authorization
	Product
}

func NewRepostitory(db *gorm.DB) *Repository {
	return &Repository{
		NewAuthPostgres(db),
		NewProductPostgres(db),
	}
}
