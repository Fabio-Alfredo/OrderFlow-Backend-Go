package user

import (
	"Auth-Service/internal/domain"
	"Auth-Service/internal/domain/models"
	"Auth-Service/internal/repository"
	"Auth-Service/pkg/config"
	"Auth-Service/pkg/logger"
	"Auth-Service/pkg/logger/console"
	"Auth-Service/pkg/obfuscate"
	"context"
	"errors"

	"gorm.io/gorm"
)

const (
	saveUserRepositoryTitle        = "save user repository: "
	findByEmailUserRepositoryTitle = "find by email user repository: "
)

type userRepository struct {
	config config.IConfig
	db     *gorm.DB
	logger logger.ILogger
}

func NewUserRepository(config config.IConfig, sqlDb *gorm.DB, logger logger.ILogger) repository.IUserRepository {
	return &userRepository{
		config: config,
		db:     sqlDb,
		logger: logger,
	}
}

func (r *userRepository) FindEmail(ctx context.Context, email string) (*models.User, error) {
	r.logger.Info(ctx, findByEmailUserRepositoryTitle+console.StartKey, "email", email)

	var user models.User
	err := r.db.
		Where("email = ?", email).
		First(&user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			r.logger.Error(ctx, findByEmailUserRepositoryTitle+console.ErrorKey, console.ErrorKey, "user not found")
			return nil, domain.ErrUserNotFound
		}
		r.logger.Error(ctx, findByEmailUserRepositoryTitle+console.ErrorKey, console.ErrorKey, err)
		return nil, err
	}

	r.logger.Info(ctx, findByEmailUserRepositoryTitle+console.EndKey, console.ResponseKey, obfuscate.Password(user))
	return &user, nil
}

func (r *userRepository) Save(ctx context.Context, data *models.User) error {
	r.logger.Info(ctx, saveUserRepositoryTitle+console.StartKey, console.DataKey, obfuscate.Password(*data))

	res := r.db.
		Create(&data)

	if err := res.Error; err != nil {
		r.logger.Error(ctx, saveUserRepositoryTitle+console.ErrorKey, console.ErrorKey, err)
		return err
	}

	r.logger.Info(ctx, saveUserRepositoryTitle+console.EndKey, console.RowsAffected, res.RowsAffected)
	return nil
}
