package token

import (
	"Auth-Service/internal/domain/models"
	"Auth-Service/internal/parser"
	"Auth-Service/internal/repository"
	"Auth-Service/pkg/logger"
	"Auth-Service/pkg/logger/console"
	"Auth-Service/pkg/obfuscate"
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

func (r *tokenRepository) Save(ctx context.Context, tokenData *models.Token) error {
	r.logger.Info(ctx, tokenRepositoryTitle+console.StartKey)

	res := r.db.
		Save(&tokenData)

	if res.Error != nil {
		r.logger.Error(ctx, tokenRepositoryTitle+console.ErrorKey, console.ErrorKey, res.Error)
		return res.Error
	}

	return nil
}

func (r *tokenRepository) FindByUserAndActive(ctx context.Context, userId string, active bool, tokenString string) (*models.Token, error) {
	r.logger.Info(ctx, tokenRepositoryTitle+console.StartKey)

	var tokenData models.Token
	err := r.db.
		Where("user_id = ? AND is_active = ? AND token = ?", userId, active, tokenString).
		Take(&tokenData).Error

	if err != nil {
		r.logger.Error(ctx, tokenRepositoryTitle+console.ErrorKey, console.ErrorKey, err)
		return nil, err
	}

	r.logger.Info(ctx, tokenRepositoryTitle+console.EndKey, console.ResponseKey, obfuscate.TokenAuth(tokenData))
	return &tokenData, nil
}
