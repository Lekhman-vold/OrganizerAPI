package service

import (
	org "github.com/Lekhman-vold/OrganizerApi"
	"github.com/Lekhman-vold/OrganizerApi/pkg/repository"
)

type Authorization interface {
	CreateUser(user org.User) (int error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
