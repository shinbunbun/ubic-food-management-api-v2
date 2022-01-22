package user

import (
	"encoding/json"
	"fmt"
	"hello-world/dynamodb"
	"hello-world/response"
	"hello-world/token"

	"github.com/aws/aws-lambda-go/events"
)

func UserGet(request events.APIGatewayProxyRequest, idTokenPayload token.Payload) events.APIGatewayProxyResponse {

	var userData User

	userData.UserID = idTokenPayload.Sub
	userData.Name = idTokenPayload.Name

	dynamodb.CreateTable()

	transactionUserCols, err := dynamodb.GetByDataDataType(userData.UserID, "transaction-user")
	if err != nil {
		return response.StatusCode500(err)
	}
	fmt.Printf("transactionUserCols: %+v\n", transactionUserCols)

	for _, v := range transactionUserCols {
		transaction := Transaction{}
		transaction.ID = v.ID

		transactionDateCol, err := dynamodb.GetByIDDataType(transaction.ID, "transaction-date")
		if err != nil {
			return response.StatusCode500(err)
		}
		transaction.Date = transactionDateCol.Data

		foodIDCol, err := dynamodb.GetByIDDataType(transaction.ID, "transaction-food")
		if err != nil {
			return response.StatusCode500(err)
		}
		foodID := foodIDCol.Data

		foodData, err := getFoodDataByID(foodID)
		if err != nil {
			return response.StatusCode500(err)
		}
		transaction.Food = foodData

		userData.Transactions = append(userData.Transactions, transaction)
	}

	resBody, err := json.Marshal(userData)
	if err != nil {
		return response.StatusCode500(err)
	}

	return response.StatusCode200(string(resBody))
}

func getFoodDataByID(foodId string) (Food, error) {
	foodData := Food{}
	foodData.ID = foodId

	foodNameCol, err := dynamodb.GetByIDDataType(foodData.ID, "food-name")
	if err != nil {
		return Food{}, err
	}
	foodData.Name = foodNameCol.Data

	foodMakerCol, err := dynamodb.GetByIDDataType(foodData.ID, "food-maker")
	if err != nil {
		return Food{}, err
	}
	foodData.Maker = foodMakerCol.Data

	foodImageCol, err := dynamodb.GetByIDDataType(foodData.ID, "food-image")
	if err != nil {
		return Food{}, err
	}
	foodData.ImageUrl = foodImageCol.Data

	foodStockCol, err := dynamodb.GetByIDDataType(foodData.ID, "food-stock")
	if err != nil {
		return Food{}, err
	}
	foodData.Stock = foodStockCol.IntData

	return foodData, nil
}
