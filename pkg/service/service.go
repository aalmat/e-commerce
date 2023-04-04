package service

import (
	"github.com/aalmat/e-commerce/models"
	"github.com/aalmat/e-commerce/pkg/repository"
)

type (
	Authorization interface {
		CreateUser(user models.User) (uint, error)
		GenerateToken(username, password string) (string, error)
		ParseToken(token string) (int, error)
	}
)

type Service struct {
	Authorization
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		NewAuthService(repo.Authorization),
	}
}
