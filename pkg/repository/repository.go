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
	AddProduct(sellerId uint, product models.WareHouse) (uint, error) // product id, error
	GetAllSellerProduct(sellerId uint) ([]models.WareHouse, error)
	DeleteProduct(sellerId, productId uint) error
	UpdateProduct(sellerId, productId uint, update models.UpdateWareHouse) error
	IncreaseProductQuantity(productId, quantity uint) error
}

type Product interface {
	GetAll() ([]models.Product, error)
	GetById(productId uint) (models.Product, error)
	SearchByName(search models.Search) ([]models.Product, error)
}

type Client interface {
	AddToCart(userId uint, whId uint, quantity uint) (uint, error)
	PurchaseAll(userId uint) error
	PurchaseById(userId uint, productIds []uint) error
	ShowCartProducts(userId uint) ([]models.WareHouse, error)
	DeleteFromCart(userid uint, productId uint) error
	ChangeProductQuantity(userid uint, productId uint, quantity uint) (uint, error)
	WriteComment(userId, productId uint, commentText string) (uint, error)
	RateProduct(userId, productId uint, rate uint) (uint, error)
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
