package uuid

import "github.com/google/uuid"

type IGenerator interface {
	GenerateId() string
}

type generateUuid struct {
}

func NewGeneratorUuid() IGenerator {
	return &generateUuid{}
}

func (u *generateUuid) GenerateId() string {
	idV4 := uuid.New().String()
	return idV4
}
