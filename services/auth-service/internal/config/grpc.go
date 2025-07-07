package config

import "time"

type GRPCConfig struct {
	Port    int           `mapstructure:"port" validate:"required"`
	Timeout time.Duration `mapstructure:"timeout"`
}
