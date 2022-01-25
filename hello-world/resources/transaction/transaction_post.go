package transaction

import (
	"encoding/json"
	"hello-world/dynamodb"
	"hello-world/response"
	"hello-world/token"
	"hello-world/types"
	"time"

	"github.com/aws/aws-lambda-go/events"
)

type requestBody struct {
	FoodId string `json:"foodId"`
}

func TransactionPost(request events.APIGatewayProxyRequest, idTokenPayload token.Payload) events.APIGatewayProxyResponse {

	bodyStr := request.Body
	var body requestBody
	err := json.Unmarshal([]byte(bodyStr), &body)
	if err != nil {
		return response.StatusCode400(err)
	}

	var transaction types.Transaction
	transaction.ID, err = dynamodb.GenerateID()
	if err != nil {
		return response.StatusCode500(err)
	}
	transaction.Date = int(time.Now().Unix())
	food := types.Food{
		ID: body.FoodId,
	}

	dynamodb.CreateTable()
	err = food.Get()
	if err != nil {
		return response.StatusCode500(err)
	}
	transaction.Food = food

	err = transaction.Put(idTokenPayload.Sub)
	if err != nil {
		return response.StatusCode500(err)
	}

	resBody, err := json.Marshal(transaction)
	if err != nil {
		return response.StatusCode500(err)
	}

	return response.StatusCode200(string(resBody))
}
