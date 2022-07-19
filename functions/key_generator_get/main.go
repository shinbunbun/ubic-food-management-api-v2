package main

import (
	"encoding/json"
	"fmt"
	"ubic-food/tools/keypair"
	"ubic-food/tools/response"
	"ubic-food/tools/token"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	clientId := request.QueryStringParameters["clientId"]

	idTokenPayload, err := token.GetIdTokenPayloadByRequest(request)
	if err != nil {
		fmt.Println("Error getting id token payload: ", err.Error())
		return response.StatusCode500(err), nil
	}
	userId := idTokenPayload.Sub
	if userId != "U6b53f4ad79a23f5427119cb44f08dbd7" {
		return response.StatusCode401(nil), nil
	}

	var keyPair keypair.KeyPair
	err = keyPair.Generate()
	if err != nil {
		fmt.Println("Error generating key pair: ", err.Error())
		return response.StatusCode500(err), nil
	}

	err = keyPair.SaveToDb(clientId)
	if err != nil {
		fmt.Println("Error saving key pair to db: ", err.Error())
		return response.StatusCode500(err), nil
	}

	resBody, err := json.Marshal(keyPair)
	if err != nil {
		fmt.Println("Error marshalling key pair: ", err.Error())
		return response.StatusCode500(err), nil
	}

	return response.StatusCode200(string(resBody)), nil
}

func main() {
	lambda.Start(handler)
}
