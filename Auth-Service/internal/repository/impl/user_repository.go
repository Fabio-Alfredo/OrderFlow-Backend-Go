package impl

import (
	"Auth-Service/internal/repository/contract"
	"Auth-Service/internal/repository/entities"
	"Auth-Service/pkg/config"
	"Auth-Service/pkg/logger"
	"Auth-Service/pkg/logger/console"
	"context"
	"errors"

	"gorm.io/gorm"
)

const (
	userRepositoryTitle = "userRepository: "
)

type userRepository struct {
	config config.IConfig
	db     *gorm.DB
	logger logger.ILogger
}

func NewUserRepository(config config.IConfig, sqlDb *gorm.DB) contract.IUserRepository {
	return &userRepository{
		config: config,
		db:     sqlDb,
	}
}

func (r *userRepository) Save(ctx context.Context, data *entities.User) error {
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

func (r *userRepository) FindEmail(ctx context.Context, email string) (entities.User, error) {
	r.logger.Info(ctx, userRepositoryTitle+console.StartKey, "email", email)

	var user entities.User
	err := r.db.Table("users_tb").
		Where("email = ?", email).
		First(&user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			r.logger.Error(ctx, userRepositoryTitle+console.ErrorKey, "message", "user not found")
			return entities.User{}, errors.New("user not found")
		}
		r.logger.Error(ctx, userRepositoryTitle+console.ErrorKey, "error", err)
		return entities.User{}, err
	}

	r.logger.Info(ctx, userRepositoryTitle+console.EndKey, "user", user)
	return user, nil
}
