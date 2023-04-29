package service

import (
	"github.com/aalmat/e-commerce/models"
	"github.com/aalmat/e-commerce/pkg/repository"
)

type ClientService struct {
	repository repository.Client
}

func (c ClientService) AddToCart(productId uint, quantity uint) (uint, error) {
	return c.repository.AddToCart(productId, quantity)
}

func (c ClientService) ShowCartProducts() ([]models.Product, error) {
	return c.repository.ShowCartProducts()
}

func (c ClientService) DeleteFromCart(productId uint, quantity uint) (uint, error) {
	return c.repository.DeleteFromCart(productId, quantity)
}

func NewClientService(repo repository.Client) *ClientService {
	return &ClientService{repo}
}
