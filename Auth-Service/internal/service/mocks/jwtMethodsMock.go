package mocks

import (
	"Auth-Service/internal/domain"
	"Auth-Service/internal/service"
	"errors"
)

type jwtMethodsMock struct {
	isError bool
}

func NewJwtMethodsMock(isError bool) service.IJWTMethods {
	return &jwtMethodsMock{
		isError: isError,
	}
}

func (j *jwtMethodsMock) GenerateJWT(_ *domain.User) (string, error) {
	if j.isError {
		return "", errors.New("error")
	}
	return "token", nil
}

func (j *jwtMethodsMock) ValidateJWT(_ string) bool {
	if j.isError {
		return false
	}
	return true
}

func (j *jwtMethodsMock) GetClaims(_ string) (*domain.JWTClaims, error) {
	return nil, nil
}
