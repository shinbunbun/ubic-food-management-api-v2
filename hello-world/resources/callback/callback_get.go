package callback

import (
	"fmt"
	"hello-world/response"
	"hello-world/token"
	"net/http"
	"strings"

	"github.com/aws/aws-lambda-go/events"
)

func CallbackGet(request events.APIGatewayProxyRequest) events.APIGatewayProxyResponse {

	query := request.QueryStringParameters

	headers := http.Header{}
	for header, values := range request.MultiValueHeaders {
		for _, value := range values {
			headers.Add(header, value)
		}
	}

	requestCookie := strings.Split(headers.Get("cookie"), "; ")

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

	_, err = token.VerifyIdToken(requestCookie, idToken)
	if err != nil {
		fmt.Println("Verify Token Error:", err)
		return response.StatusCode500(err)
	}

	/* body, err := json.Marshal(idTokenPayload)
	if err != nil {
		fmt.Println("Create Body Error:", err)
		return response.StatusCode500(err)
	} */

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       idToken,
	}
}
