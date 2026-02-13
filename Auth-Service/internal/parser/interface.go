package parser

type IParser interface {
	Parser(in ...any) (any, error)
}

type IFactory interface {
	Set(key string, parser IParser) error
	Get(key string) (parser IParser, err error)
}
