package config

type IConfig interface {
	GetString(key string) string
	GetInt(key string) int
	GetBool(key string) bool
	GetStringSlice(key string) []string
	GetStringMap(key string) map[string]interface{}
}
