package repository

import (
	"github.com/aalmat/e-commerce/models"
	"github.com/jinzhu/gorm"
	"time"
)

type SellerPostgres struct {
	db *gorm.DB
}

func (p *SellerPostgres) DeleteProduct(sellerId, productId uint) error {
	//TODO implement me
	panic("implement me")
}

func (p *SellerPostgres) UpdateProduct(sellerId, productId uint, update models.UpdateWareHouse) error {
	//TODO implement me
	panic("implement me")
}

func NewSellerPostgres(db *gorm.DB) *SellerPostgres {
	return &SellerPostgres{db: db}
}

func (p *SellerPostgres) GetAll() ([]models.WareHouse, error) {
	var products []models.WareHouse
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

func (p *SellerPostgres) CreateProduct(sellerId uint, product models.Product) (uint, error) { // product id, error
	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()
	if err := p.db.Select("title", "description", "photo", "created_at", "updated_at", "quantity").Create(&product).Error; err != nil {
		return 0, err
	}
	return product.ID, nil
}
func (p *SellerPostgres) GetAllSellerProduct(sellerId uint) ([]models.WareHouse, error) {
	return nil, nil
}

func (p *SellerPostgres) GetById(sellerId uint, productId uint) (models.WareHouse, error) {
	return models.WareHouse{}, nil
}
func (p *SellerPostgres) SearchByName(keyword string) ([]models.Product, error) {
	return nil, nil
}
func (p *SellerPostgres) FilterByPrice(minPrice, maxPrice int) ([]models.WareHouse, error) {
	return nil, nil
}
func (p *SellerPostgres) FilterByRating(minRate, maxRate int) ([]models.WareHouse, error) {
	return nil, nil
}
