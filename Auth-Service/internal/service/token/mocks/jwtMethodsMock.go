package mocks

import (
	"Auth-Service/internal/domain/models"
	"Auth-Service/internal/security"
	"errors"
)

type jwtMethodsMock struct {
	isError bool
}

func NewJwtMethodsMock(isError bool) security.IJWTMethods {
	return &jwtMethodsMock{
		isError: isError,
	}
}

func (j *jwtMethodsMock) GenerateJWT(_ *models.User) (string, error) {
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

func (j *jwtMethodsMock) GetClaims(_ string) (*security.JWTClaims, error) {
	return nil, nil
}
