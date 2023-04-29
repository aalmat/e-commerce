package repository

import (
	"github.com/aalmat/e-commerce/models"
	"github.com/jinzhu/gorm"
	"time"
)

type SellerPostgres struct {
	db *gorm.DB
}

func NewSellerPostgres(db *gorm.DB) *SellerPostgres {
	return &SellerPostgres{db: db}
}

func (p *SellerPostgres) GetAll() ([]models.Product, error) {
	var products []models.Product
	if err := p.db.Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func (p *SellerPostgres) AddProduct(house models.WareHouse) (uint, error) { // product id, error
	if err := p.db.Select("product_id", "user_id", "quantity", "price").Create(&house).Error; err != nil {
		return 0, err
	}

	p.UpdateQuantity(house.ProductId, int(house.Quantity))

	return house.ID, nil

}

func (p *SellerPostgres) UpdateQuantity(productId uint, quantity int) error {
	return nil
}

func (p *SellerPostgres) CreateProduct(product models.Product) (uint, error) { // product id, error
	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()
	if err := p.db.Select("title", "description", "photo", "created_at", "updated_at", "quantity").Create(&product).Error; err != nil {
		return 0, err
	}
	return product.ID, nil
}
func (p *SellerPostgres) GetAllSellerProduct(userId uint) ([]models.Product, error) {
	return nil, nil
}
func (p *SellerPostgres) GetById(userId uint, productId uint) (models.Product, error) {
	return models.Product{}, nil
}
func (p *SellerPostgres) SearchByName(keyword string) ([]models.Product, error) {
	return nil, nil
}
func (p *SellerPostgres) FilterByPrice(minPrice, maxPrice int) ([]models.Product, error) {
	return nil, nil
}
func (p *SellerPostgres) FilterByRating(minRate, maxRate int) ([]models.Product, error) {
	return nil, nil
}
