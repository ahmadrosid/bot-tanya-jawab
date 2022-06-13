package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	Token   string `envconfig:"TOKEN"`
	Channel string `envconfig:"CHANNEL"`
}

func Get() Config {
	cfg := Config{}
	envconfig.MustProcess("", &cfg)
	return cfg
}
