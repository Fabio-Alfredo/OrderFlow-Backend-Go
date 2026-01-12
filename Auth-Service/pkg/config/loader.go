package config

import "github.com/spf13/viper"

// Load loads the configuration from the specified path
func Load(path string) (IConfig, error) {
	// Create a new Viper instance
	v := viper.New()
	// Set the file name of the configurations file
	v.SetConfigName("application")
	// Set the path to look for the configurations file
	v.SetConfigType("yaml")
	// Add the path to the config file
	v.AddConfigPath(path)
	// Enable automatic environment variable binding
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	return NewConfig(v), nil
}
