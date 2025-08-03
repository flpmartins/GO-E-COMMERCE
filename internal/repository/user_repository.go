package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	IdPermission string    `json:"id_permission"`
	CreatedAt    time.Time `json:"created_at"`
}

type UserRepository interface {
	Create(ctx context.Context, name, email, id_permission, password string) (User, error)
	GetAll(ctx context.Context) ([]User, error)
	GetByID(ctx context.Context, id uuid.UUID) (*User, error)
	Update(ctx context.Context, id uuid.UUID, name, email, id_permission, password *string) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(ctx context.Context, name, email, id_permission, password string) (User, error) {
	query := `INSERT INTO users (id, name, email, id_permission, password) VALUES ($1, $2, $3, $4, $5) RETURNING id, name, email, id_permission, created_at`
	id := uuid.New()
	var user User

	err := r.db.QueryRowContext(ctx, query, id, name, email, id_permission, password).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.IdPermission,
		&user.CreatedAt,
	)

	if err != nil {
		return user, fmt.Errorf("erro ao criar usuário: %w", err)
	}

	return user, nil
}

func (r *userRepository) GetAll(ctx context.Context) ([]User, error) {
	query := `SELECT id, name, email, id_permission, created_at FROM users`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User

	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.IdPermission, &user.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *userRepository) GetByID(ctx context.Context, id uuid.UUID) (*User, error) {
	query := `SELECT id, name, email, id_permission, created_at FROM users WHERE id = $1`

	var user User
	err := r.db.QueryRowContext(ctx, query, id).Scan(&user.ID, &user.Name, &user.Email, &user.IdPermission, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Update(ctx context.Context, id uuid.UUID, name, email, id_permission, password *string) error {
	setParts := []string{}
	args := []interface{}{}
	argPos := 1

	if name != nil {
		setParts = append(setParts, fmt.Sprintf("name = $%d", argPos))
		args = append(args, *name)
		argPos++
	}

	if email != nil {
		setParts = append(setParts, fmt.Sprintf("email = $%d", argPos))
		args = append(args, *email)
		argPos++
	}

	if id_permission != nil {
		setParts = append(setParts, fmt.Sprintf("id_permission = $%d", argPos))
		args = append(args, *id_permission)
		argPos++
	}

	if password != nil {
		setParts = append(setParts, fmt.Sprintf("password = $%d", argPos))
		args = append(args, *password)
		argPos++
	}

	if len(setParts) == 0 {
		return fmt.Errorf("nenhum campo para atualizar")
	}

	query := fmt.Sprintf("UPDATE users SET %s WHERE id = $%d", strings.Join(setParts, ", "), argPos)
	args = append(args, id)

	_, err := r.db.ExecContext(ctx, query, args...)
	return err
}

func (r *userRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
