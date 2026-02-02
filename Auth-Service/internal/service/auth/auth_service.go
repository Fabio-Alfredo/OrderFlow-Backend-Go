package auth

import (
	"Auth-Service/internal/repository"
	"Auth-Service/internal/service"
	"Auth-Service/pkg/config"
	"Auth-Service/pkg/logger"
)

type authService struct {
	config     config.IConfig
	log        logger.ILogger
	repository repository.IUserRepository
	parsers    service.IFactory
}

func NewAuthService(config config.IConfig, log logger.ILogger, repository repository.IUserRepository, parsers service.IFactory) service.IAuthService {
	return &authService{
		config:     config,
		log:        log,
		repository: repository,
		parsers:    parsers,
	}
}
