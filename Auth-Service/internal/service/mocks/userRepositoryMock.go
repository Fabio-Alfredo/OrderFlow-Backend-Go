package mocks

import (
	"Auth-Service/internal/domain"
	"Auth-Service/internal/repository"
	"context"
	"errors"
)

type userRepositoryMock struct {
	isError   bool
	existUser bool
}

func NewUserRepositoryMock(isError bool, existUser bool) repository.IUserRepository {
	return &userRepositoryMock{
		isError:   isError,
		existUser: existUser,
	}
}

func (m *userRepositoryMock) Save(_ context.Context, _ *domain.User) error {
	if m.isError {
		return errors.New("error dummy")
	}
	return nil
}

func (m *userRepositoryMock) FindEmail(ctx context.Context, email string) (*domain.User, error) {
	if m.existUser {
		return &domain.User{
			Id:       "",
			Name:     "",
			Email:    "",
			Password: "",
			Status:   "",
		}, nil
	}
	return nil, nil
}
