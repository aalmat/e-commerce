package models

import "time"

type Order struct {
	ID           uint      `json:"id"`
	UserID       uint      `json:"user_id" binding:"required" gorm:"ForeignKey:User.ID"`
	ProductID    uint      `json:"product_id" binding:"required" gorm:"ForeignKey:Product.ID"`
	Quantity     uint      `json:"quantity" binding:"required"`
	DeliveryDate time.Time `json:"delivery_date"`
}
