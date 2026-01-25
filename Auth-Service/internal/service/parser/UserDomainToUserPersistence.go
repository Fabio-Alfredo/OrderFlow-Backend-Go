package parser

import (
	"Auth-Service/internal/domain"
	"Auth-Service/internal/repository"
	"Auth-Service/internal/service"
)

const UserDomainToUserPersistence = "UserDomainToUserPersistence"

type userDomainToUserPersistence struct {
}

func NewUserDomainToUserPersistence() service.IParser {
	return &userDomainToUserPersistence{}
}

func (p *userDomainToUserPersistence) Parser(in ...any) (any, error) {
	user := in[0].(*domain.User)

	return &repository.User{
		Id:       user.Id,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}, nil
}
