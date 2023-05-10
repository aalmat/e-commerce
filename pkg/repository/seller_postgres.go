package repository

import (
	"errors"
	"fmt"
	"github.com/aalmat/e-commerce/models"
	"github.com/jinzhu/gorm"
	"time"
)

type SellerPostgres struct {
	db *gorm.DB
}

func (p *SellerPostgres) DeleteProduct(sellerId, productId uint) error {
	tx := p.db.Begin()
	//var wh models.WareHouse
	if err := tx.Where("product_id=?", productId).Delete(&models.Cart{}).Error; err != nil && !gorm.IsRecordNotFoundError(err) {
		fmt.Println(2)
		tx.Rollback()
		return err
	}
	if err := tx.Where("id=? and user_id=?", productId, sellerId).Delete(&models.WareHouse{}).Error; err != nil {
		fmt.Println(1)
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (p *SellerPostgres) UpdateProduct(sellerId uint, update models.UpdateWareHouse) error {
	tx := p.db.Begin()

	var wh models.WareHouse
	if err := tx.Where("id = ? and user_id=?", update.WhId, sellerId).First(&wh).Error; err != nil {
		tx.Rollback()
		return err
	}

	if update.Price != 0 {
		wh.Price = update.Price
		//fmt.Println(wh.Price)
	}
	if update.Quantity != 0 {
		wh.Quantity = update.Quantity
	}

	if err := tx.Save(&wh).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}
	return nil

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
	tx := p.db.Begin()
	var wh models.WareHouse
	if err := tx.Where("product_id = ? and user_id=?", house.ProductId, house.UserID).Find(&wh).Error; !gorm.IsRecordNotFoundError(err) {
		tx.Rollback()
		return 0, errors.New("you already added warehouse")
	}

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
