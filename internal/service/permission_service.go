package service

import (
	"context"
	"fmt"
	"tmp-api/internal/repository"

	"github.com/google/uuid"
)

type PermissionService interface {
	CreatePermission(ctx context.Context, name, value string) (repository.Permission, error)
	GetPermissionByID(ctx context.Context, id uuid.UUID) (*repository.Permission, error)
	GetAllPermissions(ctx context.Context) ([]repository.Permission, error)
	UpdatePermission(ctx context.Context, id uuid.UUID, name, value *string) error
	DeletePermission(ctx context.Context, id uuid.UUID) error
}

type permissionService struct {
	repository repository.PermissionRepository
}

func NewPermissionService(repository repository.PermissionRepository) PermissionService {
	return &permissionService{repository: repository}
}

func (service *permissionService) CreatePermission(ctx context.Context, name, value string) (repository.Permission, error) {
	permission, err := service.repository.Create(ctx, name, value)
	if err != nil {
		return repository.Permission{}, fmt.Errorf("erro ao criar permissão no serviço: %w", err)
	}

	return permission, nil
}

func (service *permissionService) GetPermissionByID(ctx context.Context, id uuid.UUID) (*repository.Permission, error) {
	permission, err := service.repository.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar permissão: %w", err)
	}

	return permission, nil
}

func (service *permissionService) GetAllPermissions(ctx context.Context) ([]repository.Permission, error) {
	permissions, err := service.repository.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("erro ao listar permissões: %w", err)
	}

	if len(permissions) == 0 {
		return nil, fmt.Errorf("nenhuma permissão encontrada")
	}

	return permissions, nil
}

func (service *permissionService) UpdatePermission(ctx context.Context, id uuid.UUID, name, value *string) error {
	err := service.repository.Update(ctx, id, name, value)
	if err != nil {
		return fmt.Errorf("erro ao atualizar permissão: %w", err)
	}

	return nil
}

func (service *permissionService) DeletePermission(ctx context.Context, id uuid.UUID) error {
	err := service.repository.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("erro ao deletar permissão: %w", err)
	}

	return nil
}
