package models

import "github.com/jinzhu/gorm"

type WareHouse struct {
	gorm.Model
	ProductId uint `json:"product_id" gorm:"ForeignKey:Product.ID"`
	Quantity  uint `json:"quantity" validate:"required,min=30"`
	UserID    uint `json:"user_id" gorm:"ForeignKey:User.ID"`
	Price     uint `json:"price" validate:"required,min=5"`
}

type UpdateWareHouse struct {
	WhId     uint `json:"ware_id"`
	Quantity uint `json:"quantity" validate:"required,min=30"`
	Price    uint `json:"price" validate:"min=5"`
}
