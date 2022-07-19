package main

import (
	"context"
	"errors"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(_ context.Context, request events.APIGatewayCustomAuthorizerRequest) (events.APIGatewayCustomAuthorizerResponse, error) {
	authZHeader := request.AuthorizationToken
	if authZHeader == "" {
		return events.APIGatewayCustomAuthorizerResponse{}, errors.New("Authorization header is empty")
	}

	idTokenPayload, err := verify(authZHeader)
	if err != nil {
		return generatePolicy(idTokenPayload.Sub, "Deny", request.MethodArn), nil
	}

	return generatePolicy(idTokenPayload.Sub, "Allow", request.MethodArn), nil
}

func main() {
	lambda.Start(handler)
}
