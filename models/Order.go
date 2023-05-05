package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Order struct {
	gorm.Model
	UserID       uint      `json:"user_id" validate:"required" gorm:"ForeignKey:User.ID"`
	ProductID    uint      `json:"product_id" validate:"required" gorm:"ForeignKey:Product.ID"`
	Quantity     uint      `json:"quantity" validate:"required"`
	DeliveryDate time.Time `json:"delivery_date"`
	Status       bool      `json:"status"`
}
