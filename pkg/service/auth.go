package service

import (
	org "github.com/Lekhman-vold/OrganizerApi"
	"github.com/Lekhman-vold/OrganizerApi/pkg/repository"
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user org.User) (int, error) {
	return s.repo.CreateUser(user)
}
