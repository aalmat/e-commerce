package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Cart struct {
	gorm.Model
	UserID      uint `json:"user_id"  gorm:"ForeignKey:User.ID"`
	WareHouseID uint `json:"ware_house_id" gorm:"ForeignKey:WareHouse.ID"` //warehouse id
	Quantity    uint `json:"quantity" validate:"required,min=1"`
}

type ProductQuantity struct {
	Quantity uint `json:"quantity" binding:"required"`
}

type CartInfo struct {
	CartId    uint      `json:"cart_id"`
	UserId    uint      `json:"seller_id"`
	ProductId uint      `json:"product_id"`
	Quantity  uint      `json:"quantity"`
	Price     uint      `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}
