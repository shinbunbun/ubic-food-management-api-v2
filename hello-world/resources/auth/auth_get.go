package auth

import (
	"hello-world/random"
	"hello-world/resources/config"

	"github.com/aws/aws-lambda-go/events"
)

func AuthGet() events.APIGatewayProxyResponse {
	channelId := config.GetEnv("CHANNEL_ID")
	redirectUri := config.GetEnv("REDIRECT_URI")
	stateHash, err := random.GenerateRandomString(32)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Internal Server Error: " + err.Error(),
		}
	}
	nonceHash, err := random.GenerateRandomString(32)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Internal Server Error: " + err.Error(),
		}
	}
	url := "https://access.line.me/oauth2/v2.1/authorize?response_type=code&client_id=" + channelId + "&redirect_uri=" + redirectUri + "&state=" + stateHash + "&scope=openid profile&nonce=" + nonceHash
	return events.APIGatewayProxyResponse{
		StatusCode: 302,
		Headers: map[string]string{
			"Location": url,
		},
	}
}
