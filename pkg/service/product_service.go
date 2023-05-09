package service

import (
	"github.com/aalmat/e-commerce/models"
	"github.com/aalmat/e-commerce/pkg/repository"
)

type ProductService struct {
	repository repository.Product
}

func (p *ProductService) FilterByPrice(min, max uint) ([]models.ProductResponse, error) {
	return p.repository.FilterByPrice(min, max)
}

func (p *ProductService) ViewComment(productId uint) ([]models.CommentUser, error) {
	return p.repository.ViewComment(productId)
}

func NewProductService(r repository.Product) *ProductService {
	return &ProductService{r}
}

func (p *ProductService) GetAll() ([]models.ProductResponse, error) {
	return p.repository.GetAll()
}
func (p *ProductService) GetById(productId uint) (models.ProductResponse, error) {
	return p.repository.GetById(productId)
}

func (p *ProductService) SearchByName(search string) ([]models.ProductResponse, error) {
	return p.repository.SearchByName(search)
}

func (p *ProductService) ViewSeller(sellerId uint) ([]models.SellerResponse, error) {
	return p.repository.ViewSeller(sellerId)
}
