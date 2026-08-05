package repository

import (
	"Auth-Service/internal/domain/models"
	"context"
)

type IUserRepository interface {
	Save(ctx context.Context, data *models.User) error
	FindEmail(ctx context.Context, email string) (*models.User, error)
}

type ITokenRepository interface {
	Save(ctx context.Context, data *models.Token) error
	FindByUserAndActive(ctx context.Context, userId string, active bool, tokenString string) (*models.Token, error)
}
