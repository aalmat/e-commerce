package service

import (
	"github.com/aalmat/e-commerce/models"
	"github.com/aalmat/e-commerce/pkg/repository"
)

type SellerService struct {
	repository repository.Seller
}

func (s SellerService) AddProduct(product models.WareHouse) (uint, error) {
	return s.repository.AddProduct(product)
}

func (s SellerService) GetAllSellerProduct(userId uint) ([]models.Product, error) {
	return s.repository.GetAllSellerProduct(userId)
}

func NewSellerService(repo repository.Seller) *SellerService {
	return &SellerService{repo}
}
