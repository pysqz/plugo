package config

import ()

type Config struct {
	Plugins []string
}

var Conf *Config

func init() {
	Conf = &Config{}
}
