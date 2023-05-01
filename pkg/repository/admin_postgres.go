package repository

import (
	"fmt"
	"github.com/aalmat/e-commerce/models"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"time"
)

type AdminPostgres struct {
	db *gorm.DB
}

func (a *AdminPostgres) UpdateProduct(productId uint, update models.ProductUpdate) error {
	err := update.Validate()
	if err != nil {
		return err
	}

	var product models.Product
	if err := a.db.Where("id = ?", productId).First(&product).Error; err != nil {
		return err
	}

	if update.Title != "" {
		product.Title = update.Title
	}
	if update.Description != "" {
		product.Description = update.Description
	}
	if update.Photo != "" {
		product.Photo = update.Photo
	}

	if err := a.db.Save(&product).Error; err != nil {
		return err
	}

	return nil
}

func (a *AdminPostgres) GetProducts() ([]models.Product, error) {
	var products []models.Product
	if err := a.db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (a *AdminPostgres) CreateProduct(product models.Product) (uint, error) {
	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()
	product.Rating = 0
	product.Quantity = 0
	if err := a.db.Select("title", "description", "photo", "created_at", "updated_at", "rating", "quantity").Create(&product).Error; err != nil {
		return 0, err
	}

	return product.ID, nil
}

func (a *AdminPostgres) DeleteProduct(productId uint) error {
	if err := a.DeleteProductFromCart(productId); !gorm.IsRecordNotFoundError(err) && err != nil {
		return err
	}
	if err := a.DeleteProductFromOrder(productId); !gorm.IsRecordNotFoundError(err) && err != nil {
		return err
	}
	if err := a.DeleteProductFromWare(productId); !gorm.IsRecordNotFoundError(err) && err != nil {
		return err
	}
	if err := a.db.Where("id = ?", productId).Delete(&models.Product{}).Error; err != nil {
		return err
	}
	return nil
}

func NewAdminPostgres(db *gorm.DB) *AdminPostgres {
	return &AdminPostgres{db}
}

func (a *AdminPostgres) DeleteProductFromWare(productId uint) error {
	var wh models.WareHouse
	if err := a.db.Find(&wh, productId).Error; !gorm.IsRecordNotFoundError(err) && err != nil {
		return err
	}
	fmt.Println(wh)
	err := a.db.Delete(&wh).Error
	if !gorm.IsRecordNotFoundError(err) {
		return nil
	}
	return err
}

func (a *AdminPostgres) DeleteProductFromCart(productId uint) error {
	var wh models.Cart
	if err := a.db.Find(&wh, productId).Error; !gorm.IsRecordNotFoundError(err) && err != nil {
		return err
	}

	err := a.db.Delete(&wh).Error
	if !gorm.IsRecordNotFoundError(err) {
		return nil
	}
	return err
}

func (a *AdminPostgres) DeleteProductFromOrder(productId uint) error {
	var wh models.Order
	if err := a.db.Find(&wh, productId).Error; !gorm.IsRecordNotFoundError(err) && err != nil {
		return err
	}
	err := a.db.Delete(&wh).Error
	if !gorm.IsRecordNotFoundError(err) {
		return nil
	}
	return err
}

func (a *AdminPostgres) GetAllOrders() ([]models.Order, error) {
	var order []models.Order
	if err := a.db.Find(&order).Error; err != nil {
		return nil, err
	}

	return order, nil

}

func (a *AdminPostgres) SaveOrder(order models.Order) error {
	err := a.db.Save(&order).Error
	return err
}

func (a *AdminPostgres) CheckOrders(tickInterval time.Duration) {
	ticker := time.NewTicker(tickInterval)
	defer ticker.Stop()

	for {
		orders, err := a.GetAllOrders()
		if err != nil {
			logrus.Println("Error getting orders:", err)
			continue
		}
		select {
		case <-ticker.C:
			for _, order := range orders {
				if order.Status != true && time.Now().After(order.DeliveryDate) {
					order.Status = true
					a.SaveOrder(order)
					if err != nil {
						logrus.Println("Error updating order:", err)
						continue
					}
					logrus.Println("Order successfully delivered")
				}
			}
		}
	}
}
