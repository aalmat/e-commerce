package models

import "github.com/jinzhu/gorm"

type Cart struct {
	gorm.Model
	UserID    uint `json:"user_id" binding:"required" gorm:"ForeignKey:User.ID"`
	ProductID uint `json:"product_id" binding:"required" gorm:"ForeignKey:Product.ID"`
	Quantity  uint `json:"quantity" binding:"required"`
}
