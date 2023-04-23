package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	UserType  Role   `json:"user_type" binding:"required"`
	Email     string `json:"email" binding:"required" gorm:"not null;unique"`
	Phone     string `json:"phone" binding:"required" gorm:"not null;unique"`
	Password  string `json:"password" binding:"required"`
}

type Role int

const (
	Admin  Role = 1
	Client Role = 2
	Seller Role = 3
)
