package token

import (
	"Auth-Service/internal/domain"
	"Auth-Service/internal/parser"
	"Auth-Service/internal/repository"
	"Auth-Service/pkg/logger"
	"Auth-Service/pkg/logger/console"
	"context"

	"gorm.io/gorm"
)

const (
	tokenRepositoryTitle = "tokenRepository: "
)

type tokenRepository struct {
	db      *gorm.DB
	logger  logger.ILogger
	parsers parser.IFactory
}

func NewTokenRepository(sqlDb *gorm.DB, logger logger.ILogger, parsers parser.IFactory) repository.ITokenRepository {
	return &tokenRepository{
		db:      sqlDb,
		logger:  logger,
		parsers: parsers,
	}
}

func (r *tokenRepository) Save(ctx context.Context, data *domain.Token) (*domain.Token, error) {
	Parser, _ := r.parsers.Get(parser.TokenDomainToTokenRepositoryParser)
	r.logger.Info(ctx, tokenRepositoryTitle+console.StartKey)

	parsed, err := Parser.Parser(data)
	if err != nil {
		r.logger.Error(ctx, tokenRepositoryTitle+console.ErrorKey, console.ErrorKey, err.Error())
		return nil, err
	}

	token := parsed.(*repository.Token)

	res := r.db.
		Create(&token)

	if res.Error != nil {
		r.logger.Error(ctx, tokenRepositoryTitle+console.ErrorKey, console.ErrorKey, res.Error)
		return nil, res.Error
	}

	parserRes, _ := r.parsers.Get(parser.TokenRepositoryToTokenDomainParser)
	parsed, err = parserRes.Parser(token)
	if err != nil {
		r.logger.Error(ctx, tokenRepositoryTitle+console.ErrorKey, console.ErrorKey, err.Error())
		return nil, err
	}

	return parsed.(*domain.Token), nil
}

func (r *tokenRepository) FindAllByUserAndActive(ctx context.Context, userId string, active bool) ([]domain.Token, error) {
	r.logger.Info(ctx, tokenRepositoryTitle+console.StartKey)

	var tokens []repository.Token
	err := r.db.
		Where("user_id = ? AND is_active = ?", userId, active).
		Find(&tokens).Error

	if err != nil {
		r.logger.Error(ctx, tokenRepositoryTitle+console.ErrorKey, console.ErrorKey, err)
		return nil, err
	}

	resp, err := r.parsedTokens(tokens)
	if err != nil {
		r.logger.Error(ctx, tokenRepositoryTitle+console.ErrorKey, console.ErrorKey, err.Error())
		return nil, err
	}
	return resp, nil
}

func (r *tokenRepository) parsedTokens(tokens []repository.Token) ([]domain.Token, error) {
	var tokensDomain []domain.Token
	Parser, _ := r.parsers.Get(parser.TokenRepositoryToTokenDomainParser)

	for _, token := range tokens {
		parsed, err := Parser.Parser(&token)
		if err != nil {
			return nil, err
		}
		tokensDomain = append(tokensDomain, *parsed.(*domain.Token))
	}
	return tokensDomain, nil
}
