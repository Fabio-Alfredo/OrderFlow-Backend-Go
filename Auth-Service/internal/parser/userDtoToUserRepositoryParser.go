package parser

import (
	"Auth-Service/internal/dtos"
	"Auth-Service/internal/repository"
	"Auth-Service/internal/service"
	"Auth-Service/pkg/config"
)

const UserDtoToUserRepositoryParser = "UserDtoToUserRepositoryParser"

type userDtoToUserRepositoryParser struct {
	config config.IConfig
}

func NewUserDtoToUserRepositoryParser(config config.IConfig) service.IParser {
	return &userDtoToUserRepositoryParser{
		config: config,
	}
}

func (p *userDtoToUserRepositoryParser) Parser(in ...any) (any, error) {
	userDto := in[0].(*dtos.User)

	return &repository.User{
		Id:       userDto.Id,
		Name:     userDto.Name,
		Email:    userDto.Email,
		Password: userDto.Password,
		Status:   p.config.GetString("auth.registration.default.status"),
	}, nil
}
