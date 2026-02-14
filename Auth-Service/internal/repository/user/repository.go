package user

import (
	"Auth-Service/internal/parser"
	"Auth-Service/internal/repository"
	"Auth-Service/pkg/config"
	"Auth-Service/pkg/logger"

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
