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

func (s *SellerService) UpdateProduct(sellerId uint, update models.UpdateWareHouse) error {
	return s.repository.UpdateProduct(sellerId, update)
}

func (s *SellerService) AddProduct(product models.WareHouse) (uint, error) {
	return s.repository.AddProduct(product)
}

func (s *SellerService) GetAllSellerProduct(userId uint) ([]models.WareHouse, error) {
	return s.repository.GetAllSellerProduct(userId)
}

func (s *SellerService) IncreaseProductQuantity(productId, quantity uint) error {
	return s.repository.IncreaseProductQuantity(productId, quantity)
}

func NewSellerService(repo repository.Seller) *SellerService {
	return &SellerService{repo}
}
