package config

import "time"

type ServerConfig struct {
	Port         string
	WriteTimeout time.Duration
	ReadTimeout  time.Duration
}
