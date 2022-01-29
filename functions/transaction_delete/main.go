package main

import (
	"errors"
	"fmt"
	"strconv"
	"ubic-food/tools/dynamodb"
	"ubic-food/tools/response"

	"github.com/guregu/dynamo"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	transactionId := request.PathParameters["transactionId"]
	isStockIncrement, err := strconv.ParseBool(request.QueryStringParameters["is_stock_increment"])
	if err != nil {
		fmt.Println(err.Error())
		return response.StatusCode400(errors.New("is_stock_increment must be a boolean")), nil
	}

	dynamodb.CreateTable()

	item, err := dynamodb.GetByIDDataType(transactionId, "transaction-food")
	if err != nil {
		return response.StatusCode500(errors.New("transaction not found")), nil
	}

	keyed := []dynamo.Keyed{
		dynamo.Keys{transactionId, "transaction-date"},
		dynamo.Keys{transactionId, "transaction-food"},
		dynamo.Keys{transactionId, "transaction-user"},
	}

	err = dynamodb.BatchDelete(keyed)
	if err != nil {
		return response.StatusCode500(err), nil
	}

	if isStockIncrement {
		err = dynamodb.AddIntData(1, item.Data, "food-stock")
		if err != nil {
			fmt.Println(err.Error())
			return response.StatusCode500(errors.New("stock increment error")), nil
		}
	}

	return response.StatusCode204(), nil
}

func main() {
	lambda.Start(handler)
}
