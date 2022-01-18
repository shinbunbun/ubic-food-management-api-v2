package config

import (
	"os"

	"github.com/aws/aws-lambda-go/events"
)

func GetEnv(key string) string {
	return os.Getenv(key)
}

func GetRedirectUri(request events.APIGatewayProxyRequest) string {
	if GetEnv("STAGE") == "local" {
		return "http://localhost:3000/callback"
	}
	return ""
}
