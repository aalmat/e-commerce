package service

import (
	"github.com/aalmat/e-commerce/models"
	"github.com/aalmat/e-commerce/pkg/repository"
)

type ClientService struct {
	repository repository.Client
}

func (c *ClientService) PurchaseAll(userId uint) error {
	return c.repository.PurchaseAll(userId)
}

func (c *ClientService) PurchaseById(userId uint, productId uint) error {
	return c.repository.PurchaseById(userId, productId)
}

func (c *ClientService) RateProduct(rate models.Rating) (uint, error) {
	return c.repository.RateProduct(rate)
}

func (c *ClientService) WriteComment(comment models.Commentary) (uint, error) {
	return c.repository.WriteComment(comment)
}

func (c *ClientService) ChangeProductQuantity(userId uint, productId uint, quantity uint) (uint, error) {
	return c.repository.ChangeProductQuantity(userId, productId, quantity)
}

func (c *ClientService) AddToCart(userId, whId uint, quantity uint) (uint, error) {
	return c.repository.AddToCart(userId, whId, quantity)
}

func (c *ClientService) ShowCartProducts(userId uint) ([]models.CartInfo, error) {
	return c.repository.ShowCartProducts(userId)
}

func (c *ClientService) DeleteFromCart(userId, productId uint) error {
	return c.repository.DeleteFromCart(userId, productId)
}

func (c *ClientService) ShowOrders(userId uint) ([]models.Order, error) {
	return c.repository.ShowOrders(userId)
}

func NewClientService(repo repository.Client) *ClientService {
	return &ClientService{repo}
}
