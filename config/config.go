package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

type AppConfig struct {
	AppPort  string   `mapstructure:"APP_PORT"`
	AppEnv   string   `mapstructure:"APP_ENV"`
	Database Postgres `mapstructure:"Database"`
}

type Postgres struct {
	Host     string `mapstructure:"POSTGRES_HOST"`
	Port     string `mapstructure:"POSTGRES_PORT"`
	User     string `mapstructure:"POSTGRES_USER"`
	Password string `mapstructure:"POSTGRES_PASSWORD"`
	DbName   string `mapstructure:"POSTGRES_DB"`
}

func LoadConfig() (*AppConfig, error) {

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

	var config AppConfig
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Unable to unmarshal config: %v", err)
		return nil, err
	}

	return &config, nil
}
