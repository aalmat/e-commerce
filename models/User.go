package models

import "gorm.io/gorm"

type UserType int

const (
	ADMIN UserType = iota
	SELLER
	CLIENT
)

type User struct {
	gorm.Model
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	UserType  string `json:"user_type" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Phone     string `json:"phone" binding:"required"`
	Password  string `json:"password" binding:"required"`
}
