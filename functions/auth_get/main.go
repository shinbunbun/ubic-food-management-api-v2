package main

import (
	"os"
	"ubic-food/tools/hash"
	"ubic-food/tools/random"
	"ubic-food/tools/response"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	channelId := os.Getenv("CHANNEL_ID")
	redirectUri := os.Getenv("REDIRECT_URI")
	state, err := random.GenerateRandomString(32)
	if err != nil {
		return response.StatusCode500(err), nil
	}
	stateHash := hash.CreateSha3_256Hash(state)
	nonce, err := random.GenerateRandomString(32)
	nonceHash := hash.CreateSha3_256Hash(nonce)
	if err != nil {
		return response.StatusCode500(err), nil
	}
	url := "https://access.line.me/oauth2/v2.1/authorize?response_type=code&client_id=" + channelId + "&redirect_uri=" + redirectUri + "&state=" + stateHash + "&scope=openid profile&nonce=" + nonceHash
	return events.APIGatewayProxyResponse{
		StatusCode: 302,
		Headers: map[string]string{
			"Location": url,
		},
		MultiValueHeaders: map[string][]string{
			"Set-Cookie": {
				"nonce=" + nonce + "; HttpOnly;",
				"state=" + state + "; HttpOnly;",
			},
		},
	}, nil
}

func main() {
	lambda.Start(handler)
}
