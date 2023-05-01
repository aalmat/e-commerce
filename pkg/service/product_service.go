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

func (p ProductService) GetAll() ([]models.Product, error) {
	return p.repository.GetAll()
}
func (p *ProductService) GetById(productId uint) (models.Product, error) {
	return p.repository.GetById(productId)
}

func (p *ProductService) SearchByName(search models.Search) ([]models.Product, error) {
	return p.repository.SearchByName(search)
}
