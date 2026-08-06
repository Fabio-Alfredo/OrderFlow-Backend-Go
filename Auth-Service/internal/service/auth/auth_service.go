package auth

import (
	"Auth-Service/internal/domain"
	"Auth-Service/internal/domain/dtos"
	"Auth-Service/internal/domain/models"
	"Auth-Service/internal/repository"
	"Auth-Service/internal/security"
	"Auth-Service/internal/service"
	"Auth-Service/pkg/config"
	"Auth-Service/pkg/logger"
	"Auth-Service/pkg/logger/console"
	"Auth-Service/pkg/obfuscate"
	"Auth-Service/pkg/uuid"
	"context"
)

const (
	registerServiceTitle = "register service: "
	loginServiceTitle    = "login service: "
)

type authService struct {
	config     config.IConfig
	log        logger.ILogger
	repository repository.IUserRepository
	service    service.ITokenService
	uuid       uuid.IGenerator
	password   security.IHashMethods
}

func NewAuthService(config config.IConfig, log logger.ILogger, repository repository.IUserRepository, service service.ITokenService, uuid uuid.IGenerator, password security.IHashMethods) service.IAuthService {
	return &authService{
		config:     config,
		log:        log,
		repository: repository,
		service:    service,
		uuid:       uuid,
		password:   password,
	}
}

func (s *authService) Register(ctx context.Context, user *models.User) error {
	s.log.Info(ctx, registerServiceTitle+console.StartKey, console.RequestKey, obfuscate.Password(*user))

	existUser, err := s.repository.FindEmail(ctx, user.Email)
	if err != nil {
		return err
	}
	if existUser != nil {
		s.log.Error(ctx, registerServiceTitle, console.ErrorKey, domain.ErrAlreadyExists)
		return domain.ErrAlreadyExists
	}

	err = s.prepareUserForCreation(user)
	if err != nil {
		return err
	}

	err = s.repository.Save(ctx, user)
	if err != nil {
		s.log.Error(ctx, registerServiceTitle, console.ErrorKey, err)
		return err
	}

	s.log.Info(ctx, registerServiceTitle+console.EndKey, console.SuccessKey)
	return nil
}

func (s *authService) prepareUserForCreation(user *models.User) error {
	hashPass, err := s.password.Hash(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashPass
	user.Id = s.uuid.GenerateId()
	user.Status = s.config.GetString("auth.registration.default.status")

	return nil
}

func (s *authService) Login(ctx context.Context, authCredentials *dtos.LoginRequest) (string, error) {
	s.log.Info(ctx, loginServiceTitle+console.StartKey, console.RequestKey, obfuscate.AuthCredentials(*authCredentials))

	user, err := s.repository.FindEmail(ctx, authCredentials.Identifier)
	if err != nil {
		s.log.Error(ctx, loginServiceTitle, console.ErrorKey, err)
		return "", err
	}

	if !s.isValidCredentials(user, authCredentials) {
		s.log.Error(ctx, loginServiceTitle, console.ErrorKey, domain.ErrInvalidCredentials)
		return "", domain.ErrInvalidCredentials
	}

	token, err := s.service.Register(ctx, user)
	if err != nil {
		s.log.Error(ctx, loginServiceTitle, console.ErrorKey, err)
		return "", err
	}

	s.log.Info(ctx, loginServiceTitle+console.EndKey, console.SuccessKey)
	return token, nil
}

func (s *authService) isValidCredentials(user *models.User, credentials *dtos.LoginRequest) bool {
	return user != nil && s.password.Compare(user.Password, credentials.Password)
}
