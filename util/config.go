package util

import (
	"time"

	"github.com/spf13/viper"
)

// Config stores all configuration of the application
// The values are read by viper from a config file or environment variables.
type Config struct {
	Host                string        `mapstructure:"DB_HOST"`
	Port                string        `mapstructure:"DB_PORT"`
	Username            string        `mapstructure:"DB_USERNAME"`
	Password            string        `mapstructure:"DB_PASSWORD"`
	DBName              string        `mapstructure:"DB_NAME"`
	SSLMode             string        `mapstructure:"DB_SSLMODE"`
	PortServer          string        `mapstructure:"PORT"`
	ServerAddress       string        `mapstructure:"SERVER_ADDRESS"`
	TokenSymmetricKey   string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
}

// LoadConfig reads configuration from file or enviroment variables.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
