package service

import (
	"github.com/aalmat/e-commerce/models"
	"github.com/aalmat/e-commerce/pkg/repository"
)

type ClientService struct {
	repository repository.Client
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

func (c *ClientService) SearchByName(keyword string) ([]models.WareHouse, error) {
	return c.repository.SearchByName(keyword)
}
func (c *ClientService) FilterByPrice(minPrice, maxPrice int) ([]models.WareHouse, error) {
	return c.repository.FilterByPrice(minPrice, maxPrice)
}
func (c *ClientService) FilterByRating(minRate, maxRate int) ([]models.WareHouse, error) {
	return c.repository.FilterByRating(minRate, maxRate)
}

func NewClientService(repo repository.Client) *ClientService {
	return &ClientService{repo}
}
