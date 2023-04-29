package service

import (
	"errors"
	"github.com/aalmat/e-commerce/models"
	"github.com/aalmat/e-commerce/pkg/repository"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const TokenTime = 12 * time.Hour
const signInKey = "12j3g&ygfa1knaa1w"

type AuthService struct {
	repository repository.Authorization
}

func NewAuthService(repository repository.Authorization) *AuthService {
	return &AuthService{
		repository,
	}
}

func (a *AuthService) CreateUser(user models.User) (uint, error) {
	return a.repository.CreateUser(user)
}

func (a *AuthService) GenerateToken(email, password string) (string, error) {
	user, err := a.repository.GetUser(email, password)
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &MyClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenTime).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.ID,
		user.UserType,
	})

	return token.SignedString([]byte(signInKey))
}

func (a *AuthService) ParseToken(tokenString string) (uint, models.Role, error) {
	t, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid sign in method")
		}

		return []byte(signInKey), nil
	})

	if err != nil {
		return 0, models.Client, err
	}

	claims, ok := t.Claims.(*MyClaims)
	if !ok {
		return 0, models.Client, errors.New("invalid token claims")
	}
	return claims.UserId, claims.UserRole, nil

}

type MyClaims struct {
	jwt.StandardClaims
	UserId   uint        `json:"user_id"`
	UserRole models.Role `json:"user_role"`
}
