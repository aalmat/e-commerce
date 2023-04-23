package models

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Photo       string `json:"photo" binding:"required"`
	Price       int    `json:"price" binding:"required"`
	UserID      uint   `json:"userId" gorm:"ForeignKey:User.ID"`
	Quantity    uint   `json:"quantity"`
}
