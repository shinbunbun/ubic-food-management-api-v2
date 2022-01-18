package auth

import (
	"hello-world/config"
	"hello-world/random"
	"hello-world/response"

	"github.com/aws/aws-lambda-go/events"
)

func AuthGet(request events.APIGatewayProxyRequest) events.APIGatewayProxyResponse {
	channelId := config.GetEnv("CHANNEL_ID")
	redirectUri := config.GetRedirectUri(request)
	stateHash, err := random.GenerateRandomString(32)
	if err != nil {
		return response.StatusCode500(err)
	}
	nonceHash, err := random.GenerateRandomString(32)
	if err != nil {
		return response.StatusCode500(err)
	}
	url := "https://access.line.me/oauth2/v2.1/authorize?response_type=code&client_id=" + channelId + "&redirect_uri=" + redirectUri + "&state=" + stateHash + "&scope=openid profile&nonce=" + nonceHash
	return response.StatusCode302(url)
}
