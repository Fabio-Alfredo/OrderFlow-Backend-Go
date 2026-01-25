package exec

import (
	"Auth-Service/internal/domain"
	"Auth-Service/internal/repository"
	"Auth-Service/internal/service"
	"Auth-Service/pkg/config"
	"Auth-Service/pkg/logger"
)

type authService struct {
	config     config.IConfig
	log        logger.ILogger
	repository repository.IUserRepository
}

func NewAuthService(config config.IConfig, log logger.ILogger, repository repository.IUserRepository) service.IAuthService {
	return &authService{
		config:     config,
		log:        log,
		repository: repository,
	}
}

func (s *authService) Register(req *domain.RegisterRequest) *domain.RegisterResponse {
	return nil
}
func (s *authService) Login(req *domain.LoginRequest) *domain.LoginResponse {
	return nil
}
