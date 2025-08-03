package service

import (
	"context"
	"fmt"
	"tmp-api/internal/repository"
	"tmp-api/pkg/hashx"
)

type IAuthService interface {
	AuthLogin(ctx context.Context, email, password string) (*repository.User, error)
}

type AuthUserService struct {
	repository repository.UserRepository
}

func NewAuthService(repository repository.UserRepository) IAuthService {
	return &AuthUserService{repository: repository}
}

func (s *AuthUserService) AuthLogin(ctx context.Context, email, password string) (*repository.User, error) {
	user, err := s.repository.GetByEmail(ctx, email)

	if err != nil {
		return nil, fmt.Errorf("usuário não encontrado: %w", err)
	}

	if err := hashx.CompareHashAndPassword(user.Password, password); err != nil {
		return nil, fmt.Errorf("senha incorreta")
	}

	userCopy := *user
	userCopy.Password = ""
	return &userCopy, nil

}
