package parser

import (
	"Auth-Service/internal/domain"
	"Auth-Service/internal/repository"
)

const (
	UserRepositoryToUserDomainParser = "userRepositoryToUserDomainParser"
)

type userRepositoryToUserDomainParser struct {
}

func NewUserRepositoryToUserDomainParser() IParser {
	return &userRepositoryToUserDomainParser{}
}

func (p *userRepositoryToUserDomainParser) Parser(in ...any) (any, error) {
	user, ok := in[0].(*repository.User)

	if !ok {
		return nil, domain.ErrInvalidInput
	}

	return &domain.User{
		Id:       user.Id,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		Status:   user.Status,
	}, nil
}
