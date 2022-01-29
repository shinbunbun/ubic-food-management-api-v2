package main

import (
	"encoding/json"
	"strconv"
	"ubic-food/tools/dynamodb"
	"ubic-food/tools/response"
	"ubic-food/tools/token"
	"ubic-food/tools/types"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var userData types.User

	idTokenPayload, err := token.GetIdTokenPayloadByRequest(request)
	if err != nil {
		return response.StatusCode500(err), nil
	}

	userData.UserID = idTokenPayload.Sub
	userData.Name = idTokenPayload.Name

	transactionUserCols, err := dynamodb.GetByDataDataType(userData.UserID, "transaction-user")
	if err != nil {
		return response.StatusCode500(err), nil
	}

	for _, v := range transactionUserCols {
		transaction := types.Transaction{}
		transaction.ID = v.ID

		transactionDateCol, err := dynamodb.GetByIDDataType(transaction.ID, "transaction-date")
		if err != nil {
			return response.StatusCode500(err), nil
		}
		transaction.Date, err = strconv.Atoi(transactionDateCol.Data)
		if err != nil {
			return response.StatusCode500(err), nil
		}

		foodIDCol, err := dynamodb.GetByIDDataType(transaction.ID, "transaction-food")
		if err != nil {
			return response.StatusCode500(err), nil
		}
		foodID := foodIDCol.Data

		foodData, err := getFoodDataByID(foodID)
		if err != nil {
			return response.StatusCode500(err), nil
		}
		transaction.Food = foodData

		userData.Transactions = append(userData.Transactions, transaction)
	}

	resBody, err := json.Marshal(userData)
	if err != nil {
		return response.StatusCode500(err), nil
	}

	return response.StatusCode200(string(resBody)), nil
}

func getFoodDataByID(foodId string) (types.Food, error) {
	foodData := types.Food{}
	foodData.ID = foodId

	foodNameCol, err := dynamodb.GetByIDDataType(foodData.ID, "food-name")
	if err != nil {
		return types.Food{}, err
	}
	foodData.Name = foodNameCol.Data

	foodMakerCol, err := dynamodb.GetByIDDataType(foodData.ID, "food-maker")
	if err != nil {
		return types.Food{}, err
	}
	foodData.Maker = foodMakerCol.Data

	foodImageCol, err := dynamodb.GetByIDDataType(foodData.ID, "food-image")
	if err != nil {
		return types.Food{}, err
	}
	foodData.ImageUrl = foodImageCol.Data

	foodStockCol, err := dynamodb.GetByIDDataType(foodData.ID, "food-stock")
	if err != nil {
		return types.Food{}, err
	}
	foodData.Stock = *(foodStockCol.IntData)

	return foodData, nil
}

func main() {
	lambda.Start(handler)
}
