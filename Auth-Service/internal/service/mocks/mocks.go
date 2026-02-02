package mocks

import (
	"Auth-Service/internal/repository"
	"context"
	"errors"
	"time"
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

func (m *mockRepository) Save(ctx context.Context, data *repository.User) error {
	if m.isError {
		return errors.New("error")
	}
	return nil
}

func (m *mockRepository) FindEmail(ctx context.Context, email string) (*repository.User, error) {
	if m.existUser {
		return &repository.User{
			Id:       "",
			Name:     "",
			Email:    "",
			Password: "",
			Status:   "",
			CreateAt: time.Time{},
			UpdateAt: time.Time{},
		}, nil
	}
	return nil, nil
}
