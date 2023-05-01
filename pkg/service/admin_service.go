package service

import (
	"github.com/aalmat/e-commerce/models"
	"github.com/aalmat/e-commerce/pkg/repository"
)

type AdminService struct {
	repository repository.Admin
}

func (a *AdminService) SaveOrder(order models.Order) error {
	return a.repository.SaveOrder(order)
}

func (a *AdminService) UpdateProduct(productId uint, update models.ProductUpdate) error {
	return a.repository.UpdateProduct(productId, update)
}

func (a *AdminService) DeleteProduct(productId uint) error {
	return a.repository.DeleteProduct(productId)
}

func (a *AdminService) CreateProduct(product models.Product) (uint, error) {
	return a.repository.CreateProduct(product)
}

func NewAdminService(repo repository.Admin) *AdminService {
	return &AdminService{repo}
}

func (a *AdminService) GetAllOrders() ([]models.Order, error) {
	return a.repository.GetAllOrders()
}
