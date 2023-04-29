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
	DeleteProduct(productId uint) error
	CreateProduct(product models.Product) (uint, error)
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
