package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Cart struct {
	gorm.Model
	UserID    uint `json:"user_id" binding:"required" gorm:"ForeignKey:User.ID"`
	ProductID uint `json:"product_id" binding:"required" gorm:"ForeignKey:Product.ID"` //warehouse id
	Quantity  uint `json:"quantity" binding:"required"`
}

type ProductQuantity struct {
	Quantity uint `json:"quantity" binding:"required"`
}

type CartInfo struct {
	UserId    uint      `json:"seller_id"`
	ProductId uint      `json:"product_id"`
	Quantity  uint      `json:"quantity"`
	Price     uint      `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}
