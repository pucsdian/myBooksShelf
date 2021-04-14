package config

import (
	"time"

	"github.com/spf13/viper"
)

// Config is used to read value from .env files
type Config struct {
	Host                string        `mapstructure:"DB_HOST"`
	DBName              string        `mapstructure:"DB_NAME"`
	DBUser              string        `mapstructure:"DB_USER"`
	DBPassword          string        `mapstructure:"DB_PASSWORD"`
	Port                string        `mapstructure:"Port"`
	ServerAddress       string        `mapstructure:"SERVER_ADDRESS"`
	TokenSymmetricKey   string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
}

// LoadConfig read configuration from path having filename with extension ext
func LoadConfig(path, filename, ext string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(filename)
	viper.SetConfigType(ext)
	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}

// InitConfig loads configuration file
func InitConfig(opts ...string) (Config, error) {
	path := "."
	for _, opt := range opts {
		path = opt
	}
	return LoadConfig(path, "app", "env")
}
