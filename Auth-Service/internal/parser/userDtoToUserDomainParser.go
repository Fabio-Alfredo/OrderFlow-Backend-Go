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
	req, ok := in[0].(*dtos.RegisterRequest)
	if !ok {
		return nil, domain.ErrInvalidInput
	}

	return &domain.User{
		Id:       req.User.Id,
		Name:     req.User.Name,
		Email:    req.User.Email,
		Password: req.User.Password,
	}, nil

}
