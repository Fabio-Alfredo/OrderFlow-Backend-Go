package user

import (
	"Auth-Service/internal/repository"
	"Auth-Service/pkg/logger/console"
	"context"
)

func (r *userRepository) Save(ctx context.Context, data *repository.User) error {
	r.logger.Info(ctx, userRepositoryTitle+console.StartKey, "data", data)

	res := r.db.Table("users_tb").
		Create(&data)

	if err := res.Error; err != nil {
		r.logger.Error(ctx, userRepositoryTitle+console.ErrorKey, "error", err)
		return err
	}

	r.logger.Info(ctx, userRepositoryTitle+console.EndKey, "rowsAffected", res.RowsAffected)
	return nil
}
