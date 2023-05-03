package models

import "github.com/jinzhu/gorm"

type WareHouse struct {
	gorm.Model
	ProductId uint `json:"product_id" gorm:"ForeignKey:Product.ID"`
	Quantity  uint `json:"quantity"`
	UserID    uint `json:"user_id" gorm:"ForeignKey:User.ID"`
	Price     uint `json:"price" binding:"required"`
}

type UpdateWareHouse struct {
	WhId     uint `json:"ware_id"`
	Quantity uint `json:"quantity"`
	Price    uint `json:"price"`
}
