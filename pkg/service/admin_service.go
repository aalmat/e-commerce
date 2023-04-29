package service

import (
	"github.com/aalmat/e-commerce/models"
	"github.com/aalmat/e-commerce/pkg/repository"
)

type AdminService struct {
	repository repository.Admin
}

func (a AdminService) DeleteProduct(productId uint) error {
	return a.repository.DeleteProduct(productId)
}

func (a AdminService) CreateProduct(product models.Product) (uint, error) {
	return a.repository.CreateProduct(product)
}

func NewAdminService(repo repository.Admin) *AdminService {
	return &AdminService{repo}
}