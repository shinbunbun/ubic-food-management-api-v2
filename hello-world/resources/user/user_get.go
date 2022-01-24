package user

import (
	"encoding/json"
	"hello-world/dynamodb"
	"hello-world/response"
	"hello-world/token"
	"hello-world/types"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
)

func UserGet(request events.APIGatewayProxyRequest, idTokenPayload token.Payload) events.APIGatewayProxyResponse {

	var userData types.User

	userData.UserID = idTokenPayload.Sub
	userData.Name = idTokenPayload.Name

	dynamodb.CreateTable()

	transactionUserCols, err := dynamodb.GetByDataDataType(userData.UserID, "transaction-user")
	if err != nil {
		return response.StatusCode500(err)
	}

	for _, v := range transactionUserCols {
		transaction := types.Transaction{}
		transaction.ID = v.ID

		transactionDateCol, err := dynamodb.GetByIDDataType(transaction.ID, "transaction-date")
		if err != nil {
			return response.StatusCode500(err)
		}
		transaction.Date, err = strconv.Atoi(transactionDateCol.Data)
		if err != nil {
			return response.StatusCode500(err)
		}

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
	foodData.Stock = foodStockCol.IntData

	return foodData, nil
}
