package models

import "github.com/jinzhu/gorm"

type Rating struct {
	gorm.Model
	Rate      uint `json:"rate"`
	UserId    uint `json:"userId" gorm:"ForeignKey:User.ID"`
	ProductId uint `json:"product_id" gorm:"ForeignKey:Product.ID"`
}
