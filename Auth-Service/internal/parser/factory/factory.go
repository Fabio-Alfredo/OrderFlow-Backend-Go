package factory

import (
	"Auth-Service/internal/service"
	"errors"
	"strings"
)

type parserFactory struct {
	parsers map[string]service.IParser
}

func NewParserFactory() service.IFactory {
	return &parserFactory{
		parsers: make(map[string]service.IParser),
	}
}

func (f *parserFactory) Set(key string, parser service.IParser) error {
	if strings.TrimSpace(key) == "" {
		return errors.New("invalid the key is empty")
	}
	if parser == nil {
		return errors.New("the parser is nil")
	}

	f.parsers[key] = parser
	return nil
}

func (f *parserFactory) Get(key string) (parser service.IParser, err error) {
	if strings.TrimSpace(key) == "" {
		return nil, errors.New("invalid the key is empty")
	}

	return f.parsers[key], nil
}
