package service

import (
	"github.com/aalmat/e-commerce/models"
	"github.com/aalmat/e-commerce/pkg/repository"
)

type Authorization interface {
	CreateUser(user models.User) (uint, error)
	GenerateToken(email, password string) (string, error)
	ParseToken(token string) (uint, models.Role, error)
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
	GetById(userId uint, productId uint) (models.Product, error)
}

type Client interface {
	AddToCart(userId uint, productId uint, quantity uint) (uint, error)
	ShowCartProducts(userid uint) ([]models.WareHouse, error)
	DeleteFromCart(userid uint, productId uint) error
	ChangeProductQuantity(userid uint, productId uint, quantity uint) (uint, error)
	WriteComment(userId, productId uint, commentText string) (uint, error)
	RateProduct(userId, productId uint, rate uint) (uint, error)
	SearchByName(keyword string) ([]models.Product, error)
}

type Admin interface {
	DeleteProduct(productId uint) error
	CreateProduct(product models.Product) (uint, error)
	GetProducts() ([]models.Product, error)
	UpdateProduct(productId uint, update models.ProductUpdate) error
}

type Service struct {
	Authorization
	Product
	Seller
	Client
	Admin
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		NewAuthService(repo.Authorization),
		NewProductService(repo.Product),
		NewSellerService(repo.Seller),
		NewClientService(repo.Client),
		NewAdminService(repo.Admin),
	}
}
