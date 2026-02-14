package service

import (
	"Auth-Service/internal/domain"
	"context"
)

type IAuthService interface {
	Register(ctx context.Context, user *domain.User) (*domain.RegisterResult, error)
	Login(ctx context.Context, authCredentials *domain.AuthCredentials) (*domain.LoginResult, error)
}
