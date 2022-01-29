package transaction

import (
	"ubic-food/functions/api/dynamodb"
	"ubic-food/functions/api/response"
	"ubic-food/functions/api/token"

	"github.com/aws/aws-lambda-go/events"
	"github.com/guregu/dynamo"
)

func TransactionDelete(request events.APIGatewayProxyRequest, idTokenPayload token.Payload) events.APIGatewayProxyResponse {

	transactionId := request.PathParameters["transactionId"]

	dynamodb.CreateTable()

	keyed := []dynamo.Keyed{
		dynamo.Keys{transactionId, "transaction-date"},
		dynamo.Keys{transactionId, "transaction-food"},
		dynamo.Keys{transactionId, "transaction-user"},
	}

	err := dynamodb.BatchDelete(keyed)
	if err != nil {
		return response.StatusCode500(err)
	}

	return response.StatusCode204()
}
