package service

import (
	"github.com/aalmat/e-commerce/models"
	"github.com/aalmat/e-commerce/pkg/repository"
)

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

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	return "", nil
}

func (s *AuthService) ParseToken(token string) (int, error) {
	return 0, nil
}
