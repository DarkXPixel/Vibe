package config

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	DB    DBConfig    `mapstructure:"database"`
	GRPC  GRPCConfig  `mapstructure:"grpc"`
	Redis RedisConfig `mapstructure:"redis"`
	JWT   JWTConfig   `mapstructure:"JWT"`
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load("../.env")
	if err != nil {
		return nil, err
	}

	viper.AutomaticEnv()
	str := viper.GetString("APP_ENV")
	viper.SetConfigName(str)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../config")

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("unable to decode into struct: %w", err)
	}
	validate := validator.New()
	if err := validate.Struct(&cfg); err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			for _, fieldErr := range validationErrors {
				fmt.Printf("Field validation error: Field '%s' failed on the '%s' tag\n", fieldErr.Field(), fieldErr.Tag())
			}
		}
		return nil, fmt.Errorf("config validation failed: %w", err)
	}

	return &cfg, nil
}
