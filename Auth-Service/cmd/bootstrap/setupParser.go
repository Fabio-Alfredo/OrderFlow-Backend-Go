package bootstrap

import (
	"Auth-Service/internal/parser"
	"Auth-Service/internal/parser/factory"
	"Auth-Service/pkg/config"
)

func setupParser(config config.IConfig) parser.IFactory {
	parsers := factory.NewParserFactory()

	_ = parsers.Set(parser.UserDomainToUserRepositoryParser, parser.NewUserDomainToUserRepositoryParser(config))
	_ = parsers.Set(parser.UserDtoToUserDomainParser, parser.NewUserDtoToUserDomainParser())
	_ = parsers.Set(parser.UserRepositoryToUserDomainParser, parser.NewUserRepositoryToUserDomainParser())

	return parsers
}
