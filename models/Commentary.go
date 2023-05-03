package models

import "github.com/jinzhu/gorm"

type Commentary struct {
	gorm.Model
	UserId    uint   `json:"user_id" gorm:"ForeignKey:User.ID"`
	ProductId uint   `json:"product_id" gorm:"ForeignKey:Product.ID"`
	Text      string `json:"text"`
}

type CommentUser struct {
	Email     string `json:"email"`
	Text      string `json:"text"`
	CreatedAt string `json:"created_at"`
}
