package service

import (
	"context"
	"fmt"
	"tmp-api/internal/repository"
	hash "tmp-api/pkg/hashx"

	"github.com/google/uuid"
)

type UserService interface {
	CreateUser(ctx context.Context, name, email, id_permission, password string) (repository.User, error)
	GetUserByID(ctx context.Context, id uuid.UUID) (*repository.User, error)
	GetAllUsers(ctx context.Context) ([]repository.User, error)
	UpdateUser(ctx context.Context, id uuid.UUID, name, email, id_permission *string) error
	DeleteUser(ctx context.Context, id uuid.UUID) error
}

type userService struct {
	repository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) UserService {
	return &userService{repository: repository}
}

func (service *userService) CreateUser(ctx context.Context, name, email, id_permission, password string) (repository.User, error) {
	hashedPassword, err := hash.HashPassword(password)

	if err != nil {
		return repository.User{}, fmt.Errorf("erro ao gerar hash da senha: %w", err)
	}

	user, err := service.repository.Create(ctx, name, email, id_permission, hashedPassword)

	if err != nil {
		return repository.User{}, fmt.Errorf("erro ao criar usuário no serviço: %w", err)
	}

	return user, nil
}

func (service *userService) GetUserByID(ctx context.Context, id uuid.UUID) (*repository.User, error) {
	user, err := service.repository.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar usuário: %w", err)
	}

	return user, nil
}

func (service *userService) GetAllUsers(ctx context.Context) ([]repository.User, error) {
	users, err := service.repository.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("erro ao listar usuários: %w", err)
	}

	if len(users) == 0 {
		return nil, fmt.Errorf("nenhum usuário encontrado")
	}

	return users, nil
}

func (service *userService) UpdateUser(ctx context.Context, id uuid.UUID, name, email, id_permission *string) error {
	err := service.repository.Update(ctx, id, name, email, id_permission)
	if err != nil {
		return fmt.Errorf("erro ao atualizar usuário: %w", err)
	}

	return nil
}

func (service *userService) DeleteUser(ctx context.Context, id uuid.UUID) error {
	err := service.repository.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("erro ao deletar usuário: %w", err)
	}

	return nil
}
