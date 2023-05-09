package service

import (
	"github.com/aalmat/e-commerce/models"
	"github.com/aalmat/e-commerce/pkg/repository"
	"time"
)

type Authorization interface {
	CreateUser(user models.User) (uint, error)
	GenerateToken(email, password string) (string, error)
	ParseToken(token string) (uint, models.Role, error)
}

type Seller interface {
	AddProduct(product models.WareHouse) (uint, error) // product id, error
	GetAllSellerProduct(sellerId uint) ([]models.WareHouse, error)
	DeleteProduct(sellerId, productId uint) error
	UpdateProduct(sellerId uint, update models.UpdateWareHouse) error
	IncreaseProductQuantity(productId, quantity uint) error
}

type Product interface {
	GetAll() ([]models.ProductResponse, error)
	GetById(productId uint) (models.ProductResponse, error)
	SearchByName(search string) ([]models.ProductResponse, error)
	ViewComment(productId uint) ([]models.CommentUser, error)
	ViewSeller(sellerId uint) ([]models.SellerResponse, error)
	FilterByPrice(min, max uint) ([]models.ProductResponse, error)
}

type Client interface {
	AddToCart(userId uint, whId uint, quantity uint) (uint, error)
	ShowCartProducts(userid uint) ([]models.CartInfo, error)
	DeleteFromCart(userid uint, productId uint) error
	PurchaseAll(userId uint) error
	PurchaseById(userId uint, productId uint) error
	ChangeProductQuantity(userid uint, productId uint, quantity uint) (uint, error)
	WriteComment(comment models.Commentary) (uint, error)
	RateProduct(rate models.Rating) (uint, error)
	ShowOrders(userId uint) ([]models.Order, error)
}

type Admin interface {
	DeleteProduct(productId uint) error
	CreateProduct(product models.Product) (uint, error)
	UpdateProduct(productId uint, update models.ProductUpdate) error
	GetAllOrders() ([]models.Order, error)
	SaveOrder(order models.Order) error
	CheckOrder(tickInterval time.Duration)
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
