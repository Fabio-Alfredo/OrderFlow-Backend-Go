package user

import (
	"Auth-Service/internal/domain"
	"Auth-Service/internal/parser"
	"Auth-Service/internal/repository"
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
	config  config.IConfig
	db      *gorm.DB
	logger  logger.ILogger
	parsers parser.IFactory
}

func NewUserRepository(config config.IConfig, sqlDb *gorm.DB, logger logger.ILogger, parsers parser.IFactory) repository.IUserRepository {
	return &userRepository{
		config:  config,
		db:      sqlDb,
		logger:  logger,
		parsers: parsers,
	}
}

func (r *userRepository) FindEmail(ctx context.Context, email string) (*domain.User, error) {
	r.logger.Info(ctx, userRepositoryTitle+console.StartKey, "email", email)

	var user repository.User
	err := r.db.
		Where("email = ?", email).
		First(&user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			r.logger.Error(ctx, userRepositoryTitle+console.ErrorKey, console.ErrorKey, "user not found")
			return nil, repository.ErrUserNotFound
		}
		r.logger.Error(ctx, userRepositoryTitle+console.ErrorKey, console.ErrorKey, err)
		return nil, err
	}

	Parser, _ := r.parsers.Get(parser.UserRepositoryToUserDomainParser)
	parsed, err := Parser.Parser(&user)
	if err != nil {
		r.logger.Error(ctx, userRepositoryTitle+console.ErrorKey, console.ErrorKey, err.Error())
		return nil, err
	}

	resp := *parsed.(*domain.User)
	r.logger.Info(ctx, userRepositoryTitle+console.EndKey, console.ResponseKey, resp)
	return &resp, nil
}

func (r *userRepository) Save(ctx context.Context, domainUser *domain.User) error {
	mapper, _ := r.parsers.Get(parser.UserDomainToUserRepositoryParser)
	r.logger.Info(ctx, userRepositoryTitle+console.StartKey, console.DataKey, domainUser)

	parsed, err := mapper.Parser(domainUser)
	if err != nil {
		r.logger.Error(ctx, userRepositoryTitle+console.ErrorKey, console.ErrorKey, err.Error())
		return err
	}
	userModel := parsed.(*repository.User)

	res := r.db.
		Create(userModel)

	if err := res.Error; err != nil {
		r.logger.Error(ctx, userRepositoryTitle+console.ErrorKey, console.ErrorKey, err)
		return err
	}

	r.logger.Info(ctx, userRepositoryTitle+console.EndKey, console.RowsAffected, res.RowsAffected)
	return nil
}
