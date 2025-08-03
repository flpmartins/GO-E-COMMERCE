package service

import (
	"context"
	"fmt"
	"tmp-api/internal/repository"

	"github.com/google/uuid"
)

type UserService interface {
	CreateUser(ctx context.Context, name, email, id_permission, password string) (repository.User, error)
	GetUserByID(ctx context.Context, id uuid.UUID) (*repository.User, error)
	GetAllUsers(ctx context.Context) ([]repository.User, error)
	UpdateUser(ctx context.Context, id uuid.UUID, name, email, id_permission, password *string) error
	DeleteUser(ctx context.Context, id uuid.UUID) error
}

type userService struct {
	repository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) UserService {
	return &userService{repository: repository}
}

func (s *userService) CreateUser(ctx context.Context, name, email, id_permission, password string) (repository.User, error) {
	user, err := s.repository.Create(ctx, name, email, id_permission, password)
	if err != nil {
		return repository.User{}, fmt.Errorf("erro ao criar usuário no serviço: %w", err)
	}

	return user, nil
}

func (s *userService) GetUserByID(ctx context.Context, id uuid.UUID) (*repository.User, error) {
	user, err := s.repository.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar usuário: %w", err)
	}

	return user, nil
}

func (s *userService) GetAllUsers(ctx context.Context) ([]repository.User, error) {
	users, err := s.repository.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("erro ao listar usuários: %w", err)
	}

	if len(users) == 0 {
		return nil, fmt.Errorf("nenhum usuário encontrado")
	}

	return users, nil
}

func (s *userService) UpdateUser(ctx context.Context, id uuid.UUID, name, email, id_permission, password *string) error {
	err := s.repository.Update(ctx, id, name, email, id_permission, password)
	if err != nil {
		return fmt.Errorf("erro ao atualizar usuário: %w", err)
	}

	return nil
}

func (s *userService) DeleteUser(ctx context.Context, id uuid.UUID) error {
	err := s.repository.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("erro ao deletar usuário: %w", err)
	}

	return nil
}
