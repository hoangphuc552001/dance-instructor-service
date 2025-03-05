package config

import (
	"os"
)

type AuthConfig struct {
	InternalSecret string
}

func LoadAuthConfig() AuthConfig {
	return AuthConfig{
		InternalSecret: os.Getenv("INTERNAL_SECRET"),
	}
}
