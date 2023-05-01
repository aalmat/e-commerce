package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Order struct {
	gorm.Model
	UserID       uint      `json:"user_id" binding:"required" gorm:"ForeignKey:User.ID"`
	ProductID    uint      `json:"product_id" binding:"required" gorm:"ForeignKey:Product.ID"`
	Quantity     uint      `json:"quantity" binding:"required"`
	DeliveryDate time.Time `json:"delivery_date"`
	Status       bool      `json:"status"`
}
