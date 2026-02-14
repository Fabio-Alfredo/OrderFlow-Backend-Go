package repository

import (
	"Auth-Service/internal/domain"
	"context"
)

type IUserRepository interface {
	Save(ctx context.Context, data *domain.User) error
	FindEmail(ctx context.Context, email string) (*User, error)
}
