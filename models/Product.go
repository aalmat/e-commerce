package models

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Photo       string `json:"photo" binding:"required"`
	Quantity    uint   `json:"quantity"`
	Rating      uint   `json:"rating"`
}
