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
	GetAllSellerProduct(userId uint) ([]models.WareHouse, error)
	DeleteProduct(productId uint) error
	UpdateProduct(productId uint, update models.UpdateWareHouse) error
}

type Product interface {
	GetAll() ([]models.Product, error)
	GetById(userId uint, productId uint) (models.Product, error)
}

type Client interface {
	AddToCart(userId uint, productId uint, quantity uint) (uint, error)
	ShowCartProducts(userId uint) ([]models.WareHouse, error)
	DeleteFromCart(userid uint, productId uint) error
	ChangeProductQuantity(userid uint, productId uint, quantity uint) (uint, error)
	SearchByName(keyword string) ([]models.WareHouse, error)
	FilterByPrice(minPrice, maxPrice int) ([]models.WareHouse, error)
	FilterByRating(minRate, maxRate int) ([]models.WareHouse, error)
}

type Admin interface {
	CreateProduct(product models.Product) (uint, error)
	DeleteProduct(productId uint) error
	GetProducts() ([]models.Product, error)
	UpdateProduct(productId uint, update models.ProductUpdate) error
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
