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

func (s *AuthService) CreateUser(user models.User) (uint, error) {
	return s.repository.CreateUser(user)
}

type TokenClaim struct {
	jwt.StandardClaims
	userId uint `json:"user_id"`
}

func (s *AuthService) GenerateToken(email, password string) (string, error) {

	user, err := s.repository.GetUser(email, password)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenClaim{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenTime).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.ID,
	},
	)

	t, err := token.SignedString([]byte(signInKey))
	if err != nil {
		return "", err
	}
	return t, nil
}

func (s *AuthService) ParseToken(token string) (uint, error) {
	var userClaim TokenClaim

	t, err := jwt.ParseWithClaims(token, &userClaim, func(token *jwt.Token) (interface{}, error) {
		return []byte(signInKey), nil
	})
	if err != nil {
		return 0, err
	}

	if !t.Valid {
		return 0, errors.New("invalid token")
	}

	return userClaim.userId, nil
}
