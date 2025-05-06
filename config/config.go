package config

import (
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	AppConfig   AppConfig   `mapstructure:"AppConfig"`
	TokenConfig TokenConfig `mapstructure:"TokenConfig"`
	Database    Postgres    `mapstructure:"Database"`
}

type AppConfig struct {
	AppPort           string `mapstructure:"APP_PORT"`
	AppEnv            string `mapstructure:"APP_ENV"`
	TokenSymmetricKey string `mapstructure:"TOKEN_SYMMETRIC_KEY"`
}

type Postgres struct {
	Host     string `mapstructure:"POSTGRES_HOST"`
	Port     string `mapstructure:"POSTGRES_PORT"`
	User     string `mapstructure:"POSTGRES_USER"`
	Password string `mapstructure:"POSTGRES_PASSWORD"`
	DbName   string `mapstructure:"POSTGRES_DB"`
}

type TokenConfig struct {
	TokenSymmetricKey   string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
}

func LoadConfig() (*Config, error) {

	env := os.Getenv("ENV")
	if env == "" {
		env = "local"
	}

	viper.SetConfigName("config." + env)

	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
		return nil, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Unable to unmarshal config: %v", err)
		return nil, err
	}

	return &config, nil
}
