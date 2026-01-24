package contract

import (
	"Auth-Service/internal/repository/entities"
	"context"
)

type IUserRepository interface {
	Save(ctx context.Context, data *entities.User) error
	FindEmail(ctx context.Context, email string) (entities.User, error)
}
