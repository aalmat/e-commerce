package service

import (
	"github.com/aalmat/e-commerce/models"
	"github.com/aalmat/e-commerce/pkg/repository"
)

type ClientService struct {
	repository repository.Client
}

func (c *ClientService) RateProduct(userId, productId uint, rate uint) (uint, error) {
	return c.repository.RateProduct(userId, productId, rate)
}

func (c *ClientService) WriteComment(userId, productId uint, commentText string) (uint, error) {
	return c.repository.WriteComment(userId, productId, commentText)
}

func (c *ClientService) ChangeProductQuantity(userId uint, productId uint, quantity uint) (uint, error) {
	return c.repository.ChangeProductQuantity(userId, productId, quantity)
}

func (c *ClientService) AddToCart(userId, productId uint, quantity uint) (uint, error) {
	return c.repository.AddToCart(userId, productId, quantity)
}

func (c *ClientService) ShowCartProducts(userId uint) ([]models.WareHouse, error) {
	return c.repository.ShowCartProducts(userId)
}

func (c *ClientService) DeleteFromCart(userId, productId uint) error {
	return c.repository.DeleteFromCart(userId, productId)
}

func (c *ClientService) SearchByName(keyword string) ([]models.Product, error) {
	return c.repository.SearchByName(keyword)
}

func NewClientService(repo repository.Client) *ClientService {
	return &ClientService{repo}
}
