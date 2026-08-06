package mocks

import (
	"Auth-Service/internal/security"
	"errors"
)

type hashMock struct {
	isError bool
}

func NewHashMock(isError bool) security.IHashMethods {
	return &hashMock{
		isError: isError,
	}
}

func (h *hashMock) Hash(_ string) (string, error) {
	if h.isError {
		return "", errors.New("mock hash error")
	}
	return "hash", nil
}
func (h *hashMock) Compare(_, _ string) bool {
	return !h.isError
}
