package parser

import (
	"Auth-Service/internal/domain"
	"Auth-Service/internal/repository"
)

const TokenRepositoryToTokenDomainParser = "TokenRepositoryToTokenDomainParser"

type tokenRepositoryToTokenDomainParser struct {
}

func NewTokenRepositoryToTokenDomainParser() IParser {
	return &tokenRepositoryToTokenDomainParser{}
}

func (p *tokenRepositoryToTokenDomainParser) Parser(in ...any) (any, error) {
	token, ok := in[0].(*repository.Token)
	if !ok {
		return nil, domain.ErrInvalidInput
	}

	return &domain.Token{
		Id:        token.Id,
		UserId:    token.UserId,
		Token:     token.Token,
		IsActive:  token.IsActive,
		TimesTamp: token.TimesTamp,
	}, nil

}
