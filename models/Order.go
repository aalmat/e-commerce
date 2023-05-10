package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Order struct {
	gorm.Model
	UserID       uint      `json:"user_id" gorm:"ForeignKey:User.ID"`
	WareHouseID  uint      `json:"ware_house_id" gorm:"ForeignKey:WareHouse.ID"`
	Price        uint      `json:"price"`
	Quantity     uint      `json:"quantity"`
	DeliveryDate time.Time `json:"delivery_date"`
	Status       bool      `json:"status"`
}
