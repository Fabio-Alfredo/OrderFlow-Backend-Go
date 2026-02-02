package bootstrap

import (
	"Auth-Service/internal/parser"
	"Auth-Service/internal/parser/factory"
	"Auth-Service/internal/service"
	"Auth-Service/pkg/config"
)

func setupParser(config config.IConfig) service.IFactory {
	parsers := factory.NewParserFactory()

	_ = parsers.Set(parser.UserDtoToUserRepositoryParser, parser.NewUserDtoToUserRepositoryParser(config))

	return parsers
}
