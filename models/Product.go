package models

import (
	"errors"
	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	Title       string `json:"title" validate:"required,min=3" gorm:"not null;unique"`
	Description string `json:"description" validate:"required,min=3""`
	Photo       string `json:"photo" validate:"required"`
	Quantity    uint   `json:"quantity"`
	Rating      uint   `json:"rating"`
}

type ProductUpdate struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Photo       string `json:"photo"`
}

type ProductRating struct {
	Rating uint `json:"rating" validate:"required,min=1,max=5"`
}

func (p *ProductUpdate) Validate() error {
	if p.Title == "" && p.Description == "" && p.Photo == "" {
		return errors.New("Nothing to change")
	}
	return nil
}

type ProductResponse struct {
	Product    Product
	WareHouses []WareHouse
	Comments   []CommentUser
}
