package mocks

import (
	"Auth-Service/internal/domain/models"
	"Auth-Service/internal/service"
	"context"
	"errors"
)

type tokenServiceMock struct {
	isError bool
}

func NewTokenServiceMock(isError bool) service.ITokenService {
	return &tokenServiceMock{
		isError: isError,
	}
}

func (t *tokenServiceMock) Register(_ context.Context, _ *models.User) (string, error) {
	if t.isError {
		return "", errors.New("error")
	}

	return "token", nil
}
func (t *tokenServiceMock) IsValid(_ context.Context, _ string, _ string) (bool, error) {
	if t.isError {
		return false, errors.New("error")
	}
	return true, nil
}
