package auth

import (
	"Auth-Service/internal/domain"
	"Auth-Service/pkg/logger/console"
	"Auth-Service/pkg/obfuscate"
	"Auth-Service/pkg/utils"
	"context"

	"github.com/google/uuid"
)

const (
	registerServiceTitle = "register service: "
)

func (s *authService) Register(ctx context.Context, user *domain.User) (*domain.RegisterResult, error) {
	s.log.Info(ctx, registerServiceTitle+console.StartKey, console.RequestKey, obfuscate.RegisterService(*user))

	existUser, _ := s.repository.FindEmail(ctx, user.Email)
	if existUser != nil {
		s.log.Error(ctx, registerServiceTitle, console.ErrorKey, domain.ErrAlreadyExists)
		return nil, domain.ErrAlreadyExists
	}

	s.prepareUserForCreation(user)

	err := s.repository.Save(ctx, user)
	if err != nil {
		s.log.Error(ctx, registerServiceTitle, console.ErrorKey, err)
		return nil, err
	}

	res := s.buildResponse()
	s.log.Info(ctx, registerServiceTitle+console.EndKey, console.RequestKey, res)
	return res, nil
}

func (s *authService) prepareUserForCreation(user *domain.User) {
	hashPass, _ := utils.HashPassword(user.Password, s.config.GetInt("auth.secure.hash_cost"))
	user.Password = hashPass
	user.Id = uuid.New().String()
}

func (s *authService) buildResponse() *domain.RegisterResult {
	return &domain.RegisterResult{
		Code:    s.config.GetString("auth.register.success.code"),
		Message: s.config.GetString("auth.register.success.message"),
	}
}
