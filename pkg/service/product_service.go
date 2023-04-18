package service

import (
	"github.com/aalmat/e-commerce/models"
	"github.com/aalmat/e-commerce/pkg/repository"
)

type ProductService struct {
	repository repository.Product
}

func NewProductService(r repository.Product) *ProductService {
	return &ProductService{r}
}

func (p *ProductService) GetAll() ([]models.Product, error) {
	return p.repository.GetAll()
}
func (p *ProductService) CreateProduct(userId uint, product models.Product) (uint, error) { // product id, error
	return p.repository.CreateProduct(userId, product)
}
func (p *ProductService) GetAllSellerProduct(userId uint) ([]models.Product, error) {
	return p.repository.GetAllSellerProduct(userId)
}
func (p *ProductService) GetById(userId uint, productId uint) (models.Product, error) {
	return p.repository.GetById(userId, productId)
}
func (p *ProductService) SearchByName(keyword string) ([]models.Product, error) {
	return p.repository.SearchByName(keyword)
}
func (p *ProductService) FilterByPrice(minPrice, maxPrice int) ([]models.Product, error) {
	return p.repository.FilterByPrice(minPrice, maxPrice)
}
func (p *ProductService) FilterByRating(minRate, maxRate int) ([]models.Product, error) {
	return p.repository.FilterByRating(minRate, maxRate)
}
