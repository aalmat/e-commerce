package models

import "time"

type Cart struct {
	userId    uint `json:"user_id" gorm:"foreignKey:UserID;references:ID"`
	productId uint `json:"product_id" gorm:"foreignKey:ProductID;references:ID"`
	quantity  uint
	date      time.Time
}
