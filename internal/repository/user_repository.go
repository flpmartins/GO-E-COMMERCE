package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type UserRepository interface {
	Create(ctx context.Context, name, email string) (User, error)
	GetAll(ctx context.Context) ([]User, error)
	GetByID(ctx context.Context, id int) (*User, error)
	Update(ctx context.Context, id int, name, email string) error
	Delete(ctx context.Context, id int) error
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(ctx context.Context, name, email string) (User, error) {
	query := `INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id, name, email, created_at`

	var user User
	err := r.db.QueryRowContext(ctx, query, name, email).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.CreatedAt,
	)

	if err != nil {
		return user, fmt.Errorf("erro ao criar usuário: %w", err)
	}
	return user, nil
}

func (r *userRepository) GetAll(ctx context.Context) ([]User, error) {
	query := `SELECT id, name, email, created_at FROM users`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User

	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *userRepository) GetByID(ctx context.Context, id int) (*User, error) {
	query := `SELECT id, name, email, created_at FROM users WHERE id = $1`

	var user User
	err := r.db.QueryRowContext(ctx, query, id).Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Update(ctx context.Context, id int, name, email string) error {
	query := `UPDATE users SET name = $1, email = $2 WHERE id = $3`
	_, err := r.db.ExecContext(ctx, query, name, email, id)
	return err
}

func (r *userRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
