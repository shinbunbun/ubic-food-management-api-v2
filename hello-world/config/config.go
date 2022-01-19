package config

import (
	"os"
)

func GetEnv(key string) string {
	return os.Getenv(key)
}

func GetRedirectUri() string {
	if GetEnv("STAGE") == "local" {
		return "http://localhost:3000/callback"
	}
	return ""
}
