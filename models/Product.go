package models

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Photo       string `json:"photo" binding:"required"`
	Price       int    `json:"price" binding:"required"`
	SellerID    uint   `json:"sellerID" binding:"required"`
	Rating      int    `json:"rating"`
	RateCount   int    `json:"rateCount"`
	RateAmount  int    `json:"rateAmount"`
	User        User   `gorm:"foreignKey:SellerID"`
}
