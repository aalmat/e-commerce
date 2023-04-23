package service

import (
	"errors"
	"github.com/aalmat/e-commerce/models"
	"github.com/aalmat/e-commerce/pkg/repository"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
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

type TokenClaims struct {
	jwt.StandardClaims
	UserId   uint        `json:"user_id"`
	UserRole models.Role `json:"user_role"`
}

func (a *AuthService) GenerateToken(username, password string) (string, error) {
	//hash, err := generatePassword(password)
	//if err != nil {
	//	return "", err
	//}
	user, err := a.repository.GetUser(username, password)

	if err != nil {
		return "", err
	}
	//fmt.Println(password)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenTime).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.ID,
		user.UserType,
	})

	return token.SignedString([]byte(signInKey))

}

func (a *AuthService) ParseToken(token string) (uint, models.Role, error) {
	t, err := jwt.ParseWithClaims(token, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid sign in method")
		}

		return []byte(signInKey), nil
	})

	if err != nil {
		return 0, models.Client, err
	}

	claims, ok := t.Claims.(*TokenClaims)
	if !ok {
		return 0, models.Client, errors.New("invalid token claims")
	}
	return claims.UserId, claims.UserRole, nil
}

func generatePassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(hash), err
}
