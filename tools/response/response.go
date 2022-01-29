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

func StatusCode400(err error) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode: 400,
		Body:       "Bad Request: " + err.Error(),
	}
}

func StatusCode200(body string) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       body,
	}
}

func StatusCode204() events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode: 204,
	}
}
