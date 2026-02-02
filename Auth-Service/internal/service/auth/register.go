package auth

import (
	"Auth-Service/internal/dtos"
	"Auth-Service/internal/parser"
	"Auth-Service/internal/repository"
	"Auth-Service/internal/service"
	"Auth-Service/pkg/logger/console"
	"Auth-Service/pkg/obfuscate"
	"Auth-Service/pkg/utils"
	"context"

	"github.com/google/uuid"
)

const (
	registerServiceTitle = "register service: "
)

func (s *authService) Register(ctx context.Context, userDto *dtos.User) *service.RegisterServiceResp {
	parser, _ := s.parsers.Get(parser.UserDtoToUserRepositoryParser)
	s.log.Info(ctx, registerServiceTitle+console.StartKey, console.RequestKey, obfuscate.Register(*userDto))

	existUser, _ := s.repository.FindEmail(ctx, userDto.Email)
	if existUser != nil {
		err := &service.RegisterServiceResp{
			Code:    s.config.GetString("auth.register.errors.USER_ALREADY_EXISTS.code"),
			Message: s.config.GetString("auth.register.errors.USER_ALREADY_EXISTS.message"),
		}
		s.log.Error(ctx, registerServiceTitle, console.ErrorKey, err.Message)
		return err
	}

	s.prepareUserForCreation(userDto)
	user, _ := parser.Parser(userDto)

	err := s.repository.Save(ctx, user.(*repository.User))
	if err != nil {
		err := &service.RegisterServiceResp{
			Code:    s.config.GetString("auth.register.errors.INTERNAL.code"),
			Message: err.Error(),
		}
		s.log.Error(ctx, registerServiceTitle, console.ErrorKey, err.Message)
		return err
	}

	res := &service.RegisterServiceResp{
		Code:    s.config.GetString("auth.register.success.code"),
		Message: s.config.GetString("auth.register.success.message"),
	}
	s.log.Info(ctx, registerServiceTitle+console.EndKey, console.RequestKey, res)
	return res
}

func (s *authService) prepareUserForCreation(user *dtos.User) {
	hashPass, _ := utils.HashPassword(user.Password, s.config.GetInt("auth.secure.hash_cost"))
	user.Password = hashPass
	user.Id = uuid.New().String()
}
