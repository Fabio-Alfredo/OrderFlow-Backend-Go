package mocks

import (
	"Auth-Service/internal/domain/models"
	"Auth-Service/internal/repository"
	"context"
	"errors"
)

type tokenRepositoryMock struct {
	isError bool
}

func NewTokenRepositoryMock(isError bool) repository.ITokenRepository {
	return &tokenRepositoryMock{
		isError: isError,
	}
}

func (t *tokenRepositoryMock) Save(_ context.Context, _ *models.Token) error {
	if t.isError {
		return errors.New("error")
	}
	return nil
}
func (t *tokenRepositoryMock) FindByUserAndActive(_ context.Context, userId string, _ bool, _ string) (*models.Token, error) {
	if t.isError {
		return nil, errors.New("error")
	}
	return &models.Token{
		Id:       userId,
		Token:    "token",
		IsActive: false,
	}, nil
}
