package response

import "github.com/aws/aws-lambda-go/events"

func StatusCode500(err error) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode: 500,
		Body:       "Internal Server Error: " + err.Error(),
	}
}
