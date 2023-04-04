package repository

import (
	"github.com/aalmat/e-commerce/models"
	"github.com/jinzhu/gorm"
)

type Authorization interface {
	CreateUser(user models.User) (uint, error)
	GetUser(email, password string) (models.User, error)
}

type Repository struct {
	Authorization
}

func NewRepostitory(db *gorm.DB) *Repository {
	return &Repository{
		NewAuthPostgres(db),
	}
}
