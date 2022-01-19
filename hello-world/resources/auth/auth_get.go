package auth

import (
	"hello-world/config"
	"hello-world/hash"
	"hello-world/random"
	"hello-world/response"

	"github.com/aws/aws-lambda-go/events"
)

func AuthGet(request events.APIGatewayProxyRequest) events.APIGatewayProxyResponse {
	channelId := config.GetEnv("CHANNEL_ID")
	redirectUri := config.GetRedirectUri()
	state, err := random.GenerateRandomString(32)
	if err != nil {
		return response.StatusCode500(err)
	}
	stateHash := hash.CreateSha3_256Hash(state)
	nonce, err := random.GenerateRandomString(32)
	nonceHash := hash.CreateSha3_256Hash(nonce)
	if err != nil {
		return response.StatusCode500(err)
	}
	url := "https://access.line.me/oauth2/v2.1/authorize?response_type=code&client_id=" + channelId + "&redirect_uri=" + redirectUri + "&state=" + stateHash + "&scope=openid profile&nonce=" + nonceHash
	return events.APIGatewayProxyResponse{
		StatusCode: 302,
		Headers: map[string]string{
			"Location": url,
		},
		MultiValueHeaders: map[string][]string{
			"Set-Cookie": {
				"nonce=" + nonce,
				"state=" + state,
			},
		},
	}
}
