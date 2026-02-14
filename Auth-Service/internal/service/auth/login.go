package auth

import (
	"Auth-Service/internal/domain"
	"context"
)

func (s *authService) Login(ctx context.Context, authCredentials *domain.AuthCredentials) (*domain.LoginResult, error) {
	return nil, nil
}
