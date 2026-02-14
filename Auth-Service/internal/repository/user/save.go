package user

import (
	"Auth-Service/internal/domain"
	"Auth-Service/internal/parser"
	"Auth-Service/internal/repository"
	"Auth-Service/pkg/logger/console"
	"context"
)

func (r *userRepository) Save(ctx context.Context, domainUser *domain.User) error {
	mapper, _ := r.parsers.Get(parser.UserDomainToUserRepositoryParser)
	r.logger.Info(ctx, userRepositoryTitle+console.StartKey, "data", domainUser)

	parsed, _ := mapper.Parser(domainUser)
	userModel := parsed.(*repository.User)

	res := r.db.Table("users_tb").
		Create(userModel)

	if err := res.Error; err != nil {
		r.logger.Error(ctx, userRepositoryTitle+console.ErrorKey, "error", err)
		return err
	}

	r.logger.Info(ctx, userRepositoryTitle+console.EndKey, "rowsAffected", res.RowsAffected)
	return nil
}
