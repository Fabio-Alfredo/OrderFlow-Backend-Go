package parser

import (
	"Auth-Service/internal/domain"
	"Auth-Service/internal/repository"
)

const TokenDomainToTokenRepositoryParser = "TokenDomainToTokenRepositoryParser"

type tokenDomainToTokenRepositoryParser struct {
}

func NewTokenDomainToTokenRepositoryParser() IParser {
	return &tokenDomainToTokenRepositoryParser{}
}

func (p *tokenDomainToTokenRepositoryParser) Parser(in ...any) (any, error) {
	token, ok := in[0].(*domain.Token)

	if !ok {
		return nil, domain.ErrInvalidInput
	}

	return &repository.Token{
		Id:        token.Id,
		UserId:    token.UserId,
		Token:     token.Token,
		IsActive:  token.IsActive,
		TimesTamp: token.TimesTamp,
	}, nil
}
