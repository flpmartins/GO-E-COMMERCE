package service

import (
	"context"
	"tmp-api/internal/repository"
)

// UserService define a interface do serviço
type UserService interface {
	CreateUser(ctx context.Context, name, email string) (int, error)
	// Adicione outros métodos necessários
}

type userService struct {
	repo repository.UserRepository
}

// NewUserService retorna uma implementação concreta
func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

// Implemente os métodos da interface aqui
func (s *userService) CreateUser(ctx context.Context, name, email string) (int, error) {
	// Implementação
	return 0, nil
}
