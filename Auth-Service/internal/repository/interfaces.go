package repository

import (
	"Auth-Service/internal/domain"
	"context"
)

type IUserRepository interface {
	Save(ctx context.Context, data *domain.User) error
	FindEmail(ctx context.Context, email string) (*domain.User, error)
}

type ITokenRepository interface {
	Save(ctx context.Context, data *domain.Token) error
	FindByUserAndActive(ctx context.Context, userId string, active bool, tokenString string) (*domain.Token, error)
}
