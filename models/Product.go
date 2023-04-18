package models

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	Name   string `json:"name"`
	Photo  string `json:"photo"`
	Price  int    `json:"price"`
	Rating int    `json:"rating"`
}
