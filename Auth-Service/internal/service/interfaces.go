package service

import (
	"Auth-Service/internal/domain"
	"context"
)

type IAuthService interface {
	Register(ctx context.Context, user *domain.User) (*domain.RegisterResult, error)
	Login(ctx context.Context, authCredentials *domain.AuthCredentials) (*domain.LoginResult, error)
}

type IJWTMethods interface {
	GenerateJWT(user *domain.User) (string, error)
	ValidateJWT(token string) bool
	GetClaims(tokenString string) (*domain.JWTClaims, error)
}

type ITokenService interface {
	Register(ctx context.Context, user *domain.User) (string, error)
	IsValid(ctx context.Context, tokenString string, userId string) (bool, error)
}
