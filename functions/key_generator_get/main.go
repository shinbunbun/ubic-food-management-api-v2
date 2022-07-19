package main

import (
	"encoding/json"
	"ubic-food/tools/keypair"
	"ubic-food/tools/response"
	"ubic-food/tools/token"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	clientId := request.PathParameters["transactionId"]

	idTokenPayload, err := token.GetIdTokenPayloadByRequest(request)
	if err != nil {
		return response.StatusCode500(err), nil
	}
	userId := idTokenPayload.Sub
	if userId != "U6b53f4ad79a23f5427119cb44f08dbd7" {
		return response.StatusCode401(nil), nil
	}

	var keyPair keypair.KeyPair
	err = keyPair.Generate()
	if err != nil {
		return response.StatusCode500(err), nil
	}

	err = keyPair.SaveToDb(clientId)
	if err != nil {
		return response.StatusCode500(err), nil
	}

	resBody, err := json.Marshal(keyPair)
	if err != nil {
		return response.StatusCode500(err), nil
	}

	return response.StatusCode200(string(resBody)), nil
}

func main() {
	lambda.Start(handler)
}
