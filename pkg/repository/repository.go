package repository

import (
	"github.com/aalmat/e-commerce/models"
	"github.com/jinzhu/gorm"
)

type Authorization interface {
	CreateUser(user models.User) (uint, error)
	GetUser(email, password string) (models.User, error)
}

type Seller interface {
	AddProduct(product models.WareHouse) (uint, error) // product id, error
	GetAllSellerProduct(userId uint) ([]models.Product, error)
}

type Product interface {
	GetAll() ([]models.Product, error)
	GetById(userId uint, productId uint) (models.Product, error)
	SearchByName(keyword string) ([]models.Product, error)
	FilterByPrice(minPrice, maxPrice int) ([]models.Product, error)
	FilterByRating(minRate, maxRate int) ([]models.Product, error)
}

type Client interface {
	AddToCart(productId uint, quantity uint) (uint, error)
	ShowCartProducts() ([]models.Product, error)
	DeleteFromCart(productId uint, quantity uint) (uint, error)
}

type Admin interface {
	CreateProduct(product models.Product) (uint, error)
	DeleteProduct(productId uint) error
}

type Repository struct {
	Authorization
	Seller
	Product
	Client
	Admin
}

func NewRepostitory(db *gorm.DB) *Repository {
	return &Repository{
		NewAuthPostgres(db),
		NewSellerPostgres(db),
		NewProductPostgres(db),
		NewClientPostgres(db),
		NewAdminPostgres(db),
	}
}
