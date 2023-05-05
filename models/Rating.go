package models

import "github.com/jinzhu/gorm"

type Rating struct {
	gorm.Model
	Rate      uint `json:"rate" validate:"required,gte=1,lte=5"`
	UserId    uint `json:"userId" gorm:"ForeignKey:User.ID"`
	ProductId uint `json:"product_id" gorm:"ForeignKey:Product.ID"`
}
