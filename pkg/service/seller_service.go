package service

import (
	"github.com/aalmat/e-commerce/models"
	"github.com/aalmat/e-commerce/pkg/repository"
)

type SellerService struct {
	repository repository.Seller
}

func (s *SellerService) DeleteProduct(sellerId, productId uint) error {
	return s.repository.DeleteProduct(sellerId, productId)
}

func (s *SellerService) UpdateProduct(sellerId, productId uint, update models.UpdateWareHouse) error {
	return s.repository.UpdateProduct(sellerId, productId, update)
}

func (s *SellerService) AddProduct(sellerId uint, product models.WareHouse) (uint, error) {
	return s.repository.AddProduct(sellerId, product)
}

func (s *SellerService) GetAllSellerProduct(userId uint) ([]models.WareHouse, error) {
	return s.repository.GetAllSellerProduct(userId)
}

func NewSellerService(repo repository.Seller) *SellerService {
	return &SellerService{repo}
}
