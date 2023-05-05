package models

import (
	"errors"
	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	Title       string `json:"title" validate:"required,min=3"`
	Description string `json:"description" validate:"required,min=3""`
	Photo       string `json:"photo" validate:"required"`
	Quantity    uint   `json:"quantity" validate:"required,gte=1"`
	Rating      uint   `json:"rating"`
}

type ProductUpdate struct {
	Title       string `json:"title" validate:"min=3"`
	Description string `json:"description" validate:"min=3""`
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
