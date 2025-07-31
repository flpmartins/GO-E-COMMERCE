package repository

import (
	"context"
	"database/sql"
)

type UserRepository interface {
	Create(ctx context.Context, name, email string) (int, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(ctx context.Context, name, email string) (int, error) {
	// Implementação aqui
	return 0, nil
}
