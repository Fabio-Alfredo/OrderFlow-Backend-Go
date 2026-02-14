package parser

import (
	"Auth-Service/internal/domain"
	"Auth-Service/internal/dtos"
)

const (
	UserDtoToUserDomainParser = "UserDtoToUserDomainParser"
)

type userDtoToUserDomainParser struct {
}

func NewUserDtoToUserDomainParser() IParser {
	return &userDtoToUserDomainParser{}
}

func (p *userDtoToUserDomainParser) Parser(in ...any) (any, error) {
	userDto, ok := in[0].(*dtos.User)
	if !ok {
		return nil, domain.ErrInvalidInput
	}

	return &domain.User{
		Id:       userDto.Id,
		Name:     userDto.Name,
		Email:    userDto.Email,
		Password: userDto.Password,
	}, nil

}
