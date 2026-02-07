package user

import (
	"Auth-Service/internal/repository"
	"Auth-Service/pkg/logger/console"
	"context"
	"errors"

	"gorm.io/gorm"
)

func (r *userRepository) FindEmail(ctx context.Context, email string) (*repository.User, error) {
	r.logger.Info(ctx, userRepositoryTitle+console.StartKey, "email", email)

	var user repository.User
	err := r.db.Table("users_tb").
		Where("email = ?", email).
		First(&user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			r.logger.Error(ctx, userRepositoryTitle+console.ErrorKey, "message", "user not found")
			return nil, repository.ErrUserNotFound
		}
		r.logger.Error(ctx, userRepositoryTitle+console.ErrorKey, "error", err)
		return nil, err
	}

	r.logger.Info(ctx, userRepositoryTitle+console.EndKey, "user", user)
	return &user, nil
}
