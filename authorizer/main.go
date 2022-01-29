package main

import (
	"context"
	"errors"

	"github.com/aws/aws-lambda-go/events"
)

func handler(ctx context.Context, request events.APIGatewayCustomAuthorizerRequest) (events.APIGatewayCustomAuthorizerResponse, error) {
	jwt := request.AuthorizationToken
	if jwt == "" {
		return events.APIGatewayCustomAuthorizerResponse{}, errors.New("Authorization header is empty")
	}
}
