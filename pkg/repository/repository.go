package repository

import (
	"github.com/aalmat/e-commerce/models"
	"github.com/jinzhu/gorm"
	"time"
)

type Authorization interface {
	CreateUser(user models.User) (uint, error)
	GetUser(email, password string) (models.User, error)
}

type Seller interface {
	AddProduct(product models.WareHouse) (uint, error) // product id, error
	GetAllSellerProduct(sellerId uint) ([]models.WareHouse, error)
	DeleteProduct(sellerId, productId uint) error
	UpdateProduct(sellerId uint, update models.UpdateWareHouse) error
	IncreaseProductQuantity(productId, quantity uint) error
}

type Product interface {
	GetAll() ([]models.Product, error)
	GetById(productId uint) (models.Product, error)
	SearchByName(search models.Search) ([]models.Product, error)
	ViewComment(productId uint) ([]models.CommentUser, error)
}

type Client interface {
	AddToCart(userId uint, whId uint, quantity uint) (uint, error)
	PurchaseAll(userId uint) error
	PurchaseById(userId uint, productId uint) error
	ShowCartProducts(userId uint) ([]models.CartInfo, error)
	DeleteFromCart(userid uint, productId uint) error
	ChangeProductQuantity(userid uint, productId uint, quantity uint) (uint, error)
	WriteComment(comment models.Commentary) (uint, error)
	RateProduct(rate models.Rating) (uint, error)
	ShowOrders(userId uint) ([]models.Order, error)
}

type Admin interface {
	CreateProduct(product models.Product) (uint, error)
	DeleteProduct(productId uint) error
	GetProducts() ([]models.Product, error)
	UpdateProduct(productId uint, update models.ProductUpdate) error
	GetAllOrders() ([]models.Order, error)
	SaveOrder(order models.Order) error
	CheckOrders(tickInterval time.Duration)
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
