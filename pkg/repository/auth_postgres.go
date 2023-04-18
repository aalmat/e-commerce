package repository

import (
	"errors"
	"github.com/aalmat/e-commerce/models"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"time"
)

const maxSizeOfCart = 15

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

	a.db.Select("created_at", "updated_at", "first_name", "last_name", "phone", "user_type", "email", "password", "cart").Create(&user)
	return user.ID, nil
}

func (a *AuthPostgres) GetUser(email, password string) (models.User, error) {
	var user models.User
	if err := a.db.Where("email = ?", email).First(&user).Error; err != nil {
		return models.User{}, errors.New("User not found")
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
