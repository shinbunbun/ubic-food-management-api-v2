package callback

import (
	"encoding/json"
	"errors"
	"fmt"
	"hello-world/response"
	"strings"

	"github.com/aws/aws-lambda-go/events"
)

func CallbackGet(request events.APIGatewayProxyRequest) events.APIGatewayProxyResponse {

	query := request.QueryStringParameters
	if val, ok := query["error"]; !ok {
		fmt.Println("Query Error:", val)
		return response.StatusCode400(errors.New(val))
	}

	requestCookie := strings.Split(";", request.Headers["Cookie"])

	err := checkState(query, requestCookie, request)
	if err != nil {
		fmt.Println("State Error:", err)
		return response.StatusCode400(err)
	}

	code := query["code"]

	tokenRes, err := getAccessToken(code)
	if err != nil {
		fmt.Println("Get Token Error:", err)
		return response.StatusCode500(err)
	}
	idToken := tokenRes.IdToken

	idTokenPayload, err := verifyIdToken(requestCookie, idToken)
	if err != nil {
		fmt.Println("Verify Token Error:", err)
		return response.StatusCode500(err)
	}

	body, err := json.Marshal(idTokenPayload)
	if err != nil {
		fmt.Println("Create Body Error:", err)
		return response.StatusCode500(err)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(body),
	}
}
