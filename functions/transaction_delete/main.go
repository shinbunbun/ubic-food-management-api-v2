package main

import (
	"ubic-food/functions/api/dynamodb"
	"ubic-food/functions/api/response"

	"github.com/guregu/dynamo"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	transactionId := request.PathParameters["transactionId"]

	dynamodb.CreateTable()

	keyed := []dynamo.Keyed{
		dynamo.Keys{transactionId, "transaction-date"},
		dynamo.Keys{transactionId, "transaction-food"},
		dynamo.Keys{transactionId, "transaction-user"},
	}

	err := dynamodb.BatchDelete(keyed)
	if err != nil {
		return response.StatusCode500(err), nil
	}

	return response.StatusCode204(), nil
}

func main() {
	lambda.Start(handler)
}
