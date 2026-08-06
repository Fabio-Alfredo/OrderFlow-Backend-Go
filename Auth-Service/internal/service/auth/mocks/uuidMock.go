package mocks

import "Auth-Service/pkg/uuid"

type uuidMock struct {
}

func NewUUIDMock() uuid.IGenerator {
	return &uuidMock{}
}

func (m *uuidMock) GenerateId() string {
	return "123455"
}
