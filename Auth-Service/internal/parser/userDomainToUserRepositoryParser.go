package parser

import (
	"Auth-Service/internal/domain"
	"Auth-Service/internal/repository"
	"Auth-Service/pkg/config"
)

const UserDomainToUserRepositoryParser = "UserDomainToUserRepositoryParser"

type userDomainToUserRepositoryParser struct {
	config config.IConfig
}

func NewUserDomainToUserRepositoryParser(config config.IConfig) IParser {
	return &userDomainToUserRepositoryParser{
		config: config,
	}
}

func (p *userDomainToUserRepositoryParser) Parser(in ...any) (any, error) {
	userDto := in[0].(*domain.User)

	return &repository.User{
		Id:       userDto.Id,
		Name:     userDto.Name,
		Email:    userDto.Email,
		Password: userDto.Password,
		Status:   p.config.GetString("auth.registration.default.status"),
	}, nil
}
