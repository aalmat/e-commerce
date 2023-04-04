package repository

import (
	"errors"
	"github.com/aalmat/e-commerce/models"
	"github.com/jinzhu/gorm"
	"time"
)

type AuthPostgres struct {
	db *gorm.DB
}

func NewAuthPostgres(db *gorm.DB) *AuthPostgres {
	return &AuthPostgres{
		db,
	}
}

func (a *AuthPostgres) CreateUser(user models.User) (uint, error) {
	var err error
	user.Password, err = hashingPassword(user.Password)
	if err != nil {
		return 0, err
	}

	if err := a.db.First(&user, "email=$1", user.Email).Error; err == nil {
		return 0, errors.New("email already registered")
	}

	user.CreatedAt, user.UpdatedAt = time.Now(), time.Now()

	a.db.Select("created_at", "updated_at", "first_name", "last_name", "user_type", "phone", "email", "password").Create(&user)
	return user.ID, nil
}

func (a *AuthPostgres) GetUser(email, password string) (models.User, error) {
	return models.User{}, nil
}
