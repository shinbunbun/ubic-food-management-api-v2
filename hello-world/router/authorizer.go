package router

import (
	"strings"

	"hello-world/token"

	"github.com/aws/aws-lambda-go/events"
)

func authorizer(request events.APIGatewayProxyRequest) (token.Payload, error) {
	idToken := strings.Split(request.Headers["Authorization"], "Bearer ")[1]
	idTokenArr := strings.Split(idToken, ".")
	err := token.VerifySignature(idTokenArr)
	if err != nil {
		return token.Payload{}, err
	}

	idTokenPayload, err := token.GetIdTokenPayload(idTokenArr)
	if err != nil {
		return token.Payload{}, err
	}

	err = token.VerifyIssuer(idTokenPayload)
	if err != nil {
		return token.Payload{}, err
	}

	err = token.VerifyAud(idTokenPayload)
	if err != nil {
		return token.Payload{}, err
	}

	err = token.VerifyExp(idTokenPayload)
	if err != nil {
		return token.Payload{}, err
	}

	return idTokenPayload, nil
}
