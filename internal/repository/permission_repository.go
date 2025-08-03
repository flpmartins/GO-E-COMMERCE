package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Permission struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Value     string    `json:"value"`
	CreatedAt time.Time `json:"created_at"`
}

type PermissionRepository interface {
	Create(ctx context.Context, name, value string) (Permission, error)
	GetAll(ctx context.Context) ([]Permission, error)
	GetByID(ctx context.Context, id uuid.UUID) (*Permission, error)
	Update(ctx context.Context, id uuid.UUID, name, value *string) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type permissionRepository struct {
	db *sql.DB
}

func NewPermissionRepository(db *sql.DB) PermissionRepository {
	return &permissionRepository{db: db}
}

func (repository *permissionRepository) Create(ctx context.Context, name, value string) (Permission, error) {
	query := `INSERT INTO permissions (id, name, value) VALUES ($1, $2, $3) RETURNING id, name, value, created_at`
	id := uuid.New()
	var permission Permission

	err := repository.db.QueryRowContext(ctx, query, id, name, value).Scan(
		&permission.ID,
		&permission.Name,
		&permission.Value,
		&permission.CreatedAt,
	)

	if err != nil {
		return permission, fmt.Errorf("erro ao criar permissão: %w", err)
	}
	return permission, nil
}

func (repository *permissionRepository) GetAll(ctx context.Context) ([]Permission, error) {
	query := `SELECT id, name, value, created_at FROM permissions`
	rows, err := repository.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var permissions []Permission

	for rows.Next() {
		var permission Permission
		if err := rows.Scan(&permission.ID, &permission.Name, &permission.Value, &permission.CreatedAt); err != nil {
			return nil, err
		}
		permissions = append(permissions, permission)
	}
	return permissions, nil
}

func (repository *permissionRepository) GetByID(ctx context.Context, id uuid.UUID) (*Permission, error) {
	query := `SELECT id, name, value, created_at FROM permissions WHERE id = $1`

	var permission Permission
	err := repository.db.QueryRowContext(ctx, query, id).Scan(&permission.ID, &permission.Name, &permission.Value, &permission.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &permission, nil
}

func (repository *permissionRepository) Update(ctx context.Context, id uuid.UUID, name, value *string) error {
	setParts := []string{}
	args := []interface{}{}
	argPos := 1

	if name != nil {
		setParts = append(setParts, fmt.Sprintf("name = $%d", argPos))
		args = append(args, *name)
		argPos++
	}

	if value != nil {
		setParts = append(setParts, fmt.Sprintf("value = $%d", argPos))
		args = append(args, *value)
		argPos++
	}

	if len(setParts) == 0 {
		return fmt.Errorf("nenhum campo para atualizar")
	}

	query := fmt.Sprintf("UPDATE permissions SET %s WHERE id = $%d", strings.Join(setParts, ", "), argPos)
	args = append(args, id)

	_, err := repository.db.ExecContext(ctx, query, args...)
	return err
}

func (repository *permissionRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM permissions WHERE id = $1`
	_, err := repository.db.ExecContext(ctx, query, id)
	return err
}
