package response

import "github.com/aws/aws-lambda-go/events"

func StatusCode500(err error) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode: 500,
		Body:       "Internal Server Error: " + err.Error(),
	}
}

func StatusCode302(url string) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode: 302,
		Headers: map[string]string{
			"Location": url,
		},
	}
}
