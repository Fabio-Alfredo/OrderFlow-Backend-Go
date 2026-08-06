package service

import (
	"Auth-Service/internal/domain/dtos"
	"Auth-Service/internal/domain/models"
	"context"
)

type IAuthService interface {
	Register(ctx context.Context, user *models.User) error
	Login(ctx context.Context, authCredentials *dtos.LoginRequest) (string, error)
}

type ITokenService interface {
	Register(ctx context.Context, user *models.User) (string, error)
	IsValid(ctx context.Context, tokenString string, userId string) (bool, error)
}
