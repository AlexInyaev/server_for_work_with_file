package config

import "github.com/caarlos0/env/v11"

type Config struct {
	Port int `env:"PORT" envDefault:"5555"`
}

func NewConfig() (*Config, error) {
	var c Config

	err := env.Parse(&c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}
