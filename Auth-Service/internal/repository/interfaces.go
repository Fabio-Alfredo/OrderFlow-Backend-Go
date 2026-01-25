package repository

import (
	"context"
)

type IUserRepository interface {
	Save(ctx context.Context, data *User) error
	FindEmail(ctx context.Context, email string) (*User, error)
}
