package models

import "github.com/jinzhu/gorm"

type Commentary struct {
	gorm.Model
	UserId    uint   `json:"userId" gorm:"ForeignKey:User.ID"`
	ProductId uint   `json:"productId" gorm:"ForeignKey:Product.ID"`
	Text      string `json:"text"`
}
