package config

import (
	"log"
	"strings"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	JWT struct {
		SecretKey             string `mapstructure:"secret_key"`
		ExpirationTimeMinutes int    `mapstructure:"expiration_time_minutes"`
	}
	Database struct {
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		DBName   string `mapstructure:"dbname"`
	}
	App struct {
		Port  int  `mapstructure:"port"`
		Debug bool `mapstructure:"debug"`
	}
	Auth0 struct {
		Domain       string `mapstructure:"domain"`
		ClientID     string `mapstructure:"client_id"`
		ClientSecret string `mapstructure:"client_secret"`
		Callback     string `mapstructure:"callback"`
	}
}

func LoadConfig(path string) (*Config, error) {
	var cfg Config

	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found: %v", err)
	}

	viper.SetConfigFile(path)
	viper.AutomaticEnv()
	viper.SetEnvPrefix("app")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
