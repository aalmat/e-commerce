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

func (p *SellerPostgres) AddProduct(sellerId uint, house models.WareHouse) (uint, error) { // product id, error
	tx := p.db.Begin()
	house.UserID = sellerId
	house.CreatedAt = time.Now()
	house.UpdatedAt = time.Now()
	if err := tx.Select("product_id", "user_id", "quantity", "price", "created_at", "updated_at").Create(&house).Error; err != nil {
		tx.Rollback()
		return 0, err
	}

	if err := p.IncreaseProductQuantity(house.ProductId, house.Quantity); err != nil {
		tx.Rollback()
		return 0, err
	}
	tx.Commit()

	return house.ID, nil

}

func (p *SellerPostgres) UpdateQuantity(productId uint, quantity int) error {
	return nil
}

func (p *SellerPostgres) GetAllSellerProduct(sellerId uint) ([]models.WareHouse, error) {
	var whs []models.WareHouse
	if err := p.db.Where("user_id = ?", sellerId).Find(&whs).Error; err != nil {
		return nil, err
	}

	return whs, nil
}

func (p *SellerPostgres) IncreaseProductQuantity(productId, quantity uint) error {
	var product models.Product
	if err := p.db.First(&product, productId).Error; err != nil {
		return err
	}
	product.Quantity += quantity

	err := p.db.Save(&product).Error
	if err != nil {
		return err
	}

	return nil

}
