package main

import (
	"ubic-food/functions/api/router"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return router.Router(request)
}

func main() {
	lambda.Start(handler)
}
