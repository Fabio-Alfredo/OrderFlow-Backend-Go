package mocks

import (
	"Auth-Service/internal/domain"
	"Auth-Service/internal/service"
	"context"
)

type serviceMock struct {
	isError bool
}

func NewServiceMock(isError bool) service.IAuthService {
	return &serviceMock{
		isError: isError,
	}
}

func (s *serviceMock) Register(ctx context.Context, user *domain.User) (*domain.RegisterResult, error) {
	if s.isError {
		return nil, domain.ErrAlreadyExists
	}
	return &domain.RegisterResult{
		Code:    "200",
		Message: "Success",
	}, nil
}
func (s *serviceMock) Login(ctx context.Context, authCredentials *domain.AuthCredentials) (*domain.LoginResult, error) {
	return nil, nil
}
