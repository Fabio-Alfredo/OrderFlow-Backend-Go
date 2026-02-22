package mocks

import (
	"Auth-Service/internal/domain"
	"Auth-Service/internal/repository"
	"context"
	"errors"
)

type mockRepository struct {
	isError   bool
	existUser bool
}

func NewMockRepository(isError bool, existUser bool) repository.IUserRepository {
	return &mockRepository{
		isError:   isError,
		existUser: existUser,
	}
}

func (m *mockRepository) Save(_ context.Context, _ *domain.User) error {
	if m.isError {
		return errors.New("error dummy")
	}
	return nil
}

func (m *mockRepository) FindEmail(ctx context.Context, email string) (*domain.User, error) {
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
