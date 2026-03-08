package auth

import (
	"Auth-Service/internal/parser"
	"Auth-Service/internal/repository"
	"Auth-Service/internal/service"
	"Auth-Service/pkg/config"
	"Auth-Service/pkg/logger"
	"Auth-Service/pkg/logger/console"
	"context"
	"strings"
)

const (
	tokenServiceTitle = "token service: "
)

type tokenService struct {
	config     config.IConfig
	log        logger.ILogger
	repository repository.ITokenRepository
	jwtMethods service.JWTMethods
	parsers    parser.IFactory
}

func NewTokenService(config config.IConfig, log logger.ILogger, repository repository.ITokenRepository, jwtMethods service.JWTMethods, parsers parser.IFactory) service.ITokenService {
	return &tokenService{
		config:     config,
		log:        log,
		repository: repository,
		jwtMethods: jwtMethods,
		parsers:    parsers,
	}
}

func (s *tokenService) Register(ctx context.Context, userId string) (string, error) {

	return "", nil
}

func (s *tokenService) IsValid(ctx context.Context, tokenString string, userId string) (bool, error) {
	s.log.Info(ctx, tokenServiceTitle+console.StartKey, console.RequestKey, userId)

	err := s.cleanTokens(ctx, userId)
	if err != nil {
		s.log.Error(ctx, tokenServiceTitle+console.ErrorKey, "Clean token error: ", err)
	}

	tokens, err := s.repository.FindAllByUserAndActive(ctx, userId, true)
	if err != nil {
		s.log.Error(ctx, tokenServiceTitle+console.ErrorKey, err)
		return false, err
	}

	for _, token := range tokens {
		if strings.EqualFold(token.Token, tokenString) {
			return true, nil
		}
	}
	return false, nil
}

func (s *tokenService) cleanTokens(ctx context.Context, userId string) error {
	tokens, err := s.repository.FindAllByUserAndActive(ctx, userId, true)
	if err != nil {
		return err
	}

	for _, token := range tokens {
		if !s.jwtMethods.ValidateJWT(token.Token) {
			token.IsActive = false
			_, e := s.repository.Save(ctx, &token)
			if e != nil {
				return e
			}
		}
	}

	return nil
}
