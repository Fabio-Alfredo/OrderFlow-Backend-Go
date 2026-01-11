package config

import "github.com/spf13/viper"

type config struct {
	viper *viper.Viper
}

func NewConfig(viper *viper.Viper) Config {
	return &config{
		viper: viper,
	}
}

func (c *config) GetString(key string) string {
	return c.viper.GetString(key)
}
func (c *config) GetInt(key string) int {
	return c.viper.GetInt(key)
}
func (c *config) GetBool(key string) bool {
	return c.viper.GetBool(key)
}
func (c *config) GetStringSlice(key string) []string {
	return c.viper.GetStringSlice(key)
}

func (c *config) GetStringMap(key string) map[string]interface{} {
	return c.viper.GetStringMap(key)
}
