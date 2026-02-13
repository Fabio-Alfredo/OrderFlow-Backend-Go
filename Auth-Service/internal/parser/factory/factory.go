package factory

import (
	"Auth-Service/internal/parser"
	"errors"
	"strings"
)

type parserFactory struct {
	parsers map[string]parser.IParser
}

func NewParserFactory() parser.IFactory {
	return &parserFactory{
		parsers: make(map[string]parser.IParser),
	}
}

func (f *parserFactory) Set(key string, parser parser.IParser) error {
	if strings.TrimSpace(key) == "" {
		return errors.New("invalid the key is empty")
	}
	if parser == nil {
		return errors.New("the parser is nil")
	}

	f.parsers[key] = parser
	return nil
}

func (f *parserFactory) Get(key string) (parser parser.IParser, err error) {
	if strings.TrimSpace(key) == "" {
		return nil, errors.New("invalid the key is empty")
	}

	return f.parsers[key], nil
}
