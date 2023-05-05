package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName string `json:"first_name" validate:"required,min=3,max=40"`
	LastName  string `json:"last_name" validate:"required,min=3,max=40"`
	UserType  Role   `json:"user_type" validate:"required"`
	Email     string `json:"email" validate:"required,email" gorm:"not null;unique"`
	Phone     string `json:"phone" validate:"required" gorm:"not null;unique"`
	Password  string `json:"password" validate:"required,min=8"`
}

type SignUser struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type Role int

const (
	Admin  Role = 1
	Client Role = 2
	Seller Role = 3
)
