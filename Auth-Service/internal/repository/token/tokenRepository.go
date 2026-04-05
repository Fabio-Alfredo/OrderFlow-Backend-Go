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

func (r *tokenRepository) Save(ctx context.Context, data *domain.Token) error {
	Parser, _ := r.parsers.Get(parser.TokenDomainToTokenRepositoryParser)
	r.logger.Info(ctx, tokenRepositoryTitle+console.StartKey)

	parsed, err := Parser.Parser(data)
	if err != nil {
		r.logger.Error(ctx, tokenRepositoryTitle+console.ErrorKey, console.ErrorKey, err.Error())
		return err
	}

	token := parsed.(*repository.Token)
	res := r.db.
		Save(&token)

	if res.Error != nil {
		r.logger.Error(ctx, tokenRepositoryTitle+console.ErrorKey, console.ErrorKey, res.Error)
		return res.Error
	}

	parserRes, _ := r.parsers.Get(parser.TokenRepositoryToTokenDomainParser)
	parsed, err = parserRes.Parser(token)
	if err != nil {
		r.logger.Error(ctx, tokenRepositoryTitle+console.ErrorKey, console.ErrorKey, err.Error())
		return err
	}

	return nil
}

func (r *tokenRepository) FindByUserAndActive(ctx context.Context, userId string, active bool, tokenString string) (*domain.Token, error) {
	r.logger.Info(ctx, tokenRepositoryTitle+console.StartKey)

	var token repository.Token
	err := r.db.
		Where("user_id = ? AND is_active = ? AND token = ?", userId, active, tokenString).
		Take(&token).Error

	if err != nil {
		r.logger.Error(ctx, tokenRepositoryTitle+console.ErrorKey, console.ErrorKey, err)
		return nil, err
	}

	Parser, _ := r.parsers.Get(parser.TokenRepositoryToTokenDomainParser)
	parsed, err := Parser.Parser(&token)
	if err != nil {
		r.logger.Error(ctx, tokenRepositoryTitle+console.ErrorKey, console.ErrorKey, err.Error())
		return nil, err
	}

	resp := parsed.(*domain.Token)

	r.logger.Info(ctx, tokenRepositoryTitle+console.EndKey, console.ResponseKey, resp)
	return resp, nil
}
