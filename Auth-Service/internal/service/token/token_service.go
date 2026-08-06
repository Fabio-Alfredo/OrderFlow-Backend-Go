package token

import (
	"Auth-Service/internal/domain/models"
	"Auth-Service/internal/repository"
	"Auth-Service/internal/security"
	"Auth-Service/internal/service"
	"Auth-Service/pkg/config"
	"Auth-Service/pkg/logger"
	"Auth-Service/pkg/logger/console"
	"context"
	"errors"

	"gorm.io/gorm"
)

const (
	tokenServiceTitle = "token service: "
)

type tokenService struct {
	config     config.IConfig
	log        logger.ILogger
	repository repository.ITokenRepository
	jwtMethods security.IJWTMethods
}

func NewTokenService(config config.IConfig, log logger.ILogger, repository repository.ITokenRepository, jwtMethods security.IJWTMethods) service.ITokenService {
	return &tokenService{
		config:     config,
		log:        log,
		repository: repository,
		jwtMethods: jwtMethods,
	}
}

func (s *tokenService) Register(ctx context.Context, user *models.User) (string, error) {
	s.log.Info(ctx, tokenServiceTitle+console.StartKey, console.RequestKey, user)

	tokenString, err := s.jwtMethods.GenerateJWT(user)
	if err != nil {
		s.log.Error(ctx, tokenServiceTitle+console.ErrorKey, "Generate JWT error: ", err)
		return "", err
	}

	token := &models.Token{
		Token:    tokenString,
		IsActive: true,
	}

	err = s.repository.Save(ctx, token)
	if err != nil {
		s.log.Error(ctx, tokenServiceTitle+console.ErrorKey, "Save token error: ", err)
		return "", err
	}

	s.log.Info(ctx, tokenServiceTitle+console.EndKey, console.ResponseKey, tokenString)
	return tokenString, nil
}

func (s *tokenService) IsValid(ctx context.Context, tokenString string, userId string) (bool, error) {
	s.log.Info(ctx, tokenServiceTitle+console.StartKey, console.RequestKey, userId)

	token, err := s.repository.FindByUserAndActive(ctx, userId, true, tokenString)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			s.log.Info(ctx, tokenServiceTitle+console.FailedKey, console.ErrorKey, "Token not found")
			return false, nil
		}
		s.log.Error(ctx, tokenServiceTitle+console.ErrorKey, err)
		return false, err
	}

	if !s.isValidToken(token) {
		s.log.Info(ctx, tokenServiceTitle+console.FailedKey, console.ErrorKey, "Token is not valid")
		s.cleanToken(ctx, token)
		return false, nil
	}

	s.log.Info(ctx, tokenServiceTitle+console.EndKey, console.ResponseKey, "Token is valid")
	return true, nil
}

func (s *tokenService) isValidToken(token *models.Token) bool {
	return token != nil && s.jwtMethods.ValidateJWT(token.Token)
}

func (s *tokenService) cleanToken(ctx context.Context, token *models.Token) {
	s.log.Info(ctx, tokenServiceTitle+console.StartKey)
	if token != nil {
		token.IsActive = false
		_ = s.repository.Save(ctx, token)
	}
	s.log.Info(ctx, tokenServiceTitle+console.EndKey)
}
