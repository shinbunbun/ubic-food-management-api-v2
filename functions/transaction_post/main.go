package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"ubic-food/tools/dynamodb"
	"ubic-food/tools/response"
	"ubic-food/tools/token"
	"ubic-food/tools/types"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type requestBody struct {
	FoodId string `json:"foodId"`
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	idTokenPayload, err := token.GetIdTokenPayloadByRequest(request)
	if err != nil {
		return response.StatusCode500(err), nil
	}

	bodyStr := request.Body
	var body requestBody
	err = json.Unmarshal([]byte(bodyStr), &body)
	if err != nil {
		return response.StatusCode400(err), nil
	}

	var transaction types.Transaction
	transaction.ID, err = dynamodb.GenerateID()
	if err != nil {
		fmt.Println("Error generating ID:", err.Error())
		return response.StatusCode500(err), nil
	}
	transaction.Date = int(time.Now().Unix())
	food := types.Food{
		ID: body.FoodId,
	}

	dynamodb.CreateTable()
	err = food.Get()
	if err != nil {
		fmt.Println("Error getting food:", err.Error())
		return response.StatusCode500(err), nil
	}
	transaction.Food = food

	if food.Stock < 1 {
		return response.StatusCode400(errors.New("food is out of stock")), nil
	}

	/* err = dynamodb.AddIntData(-1, food.ID, "food-stock")
	if err != nil {
		fmt.Println("Error adding food stock:", err.Error())
		return response.StatusCode500(err)
	} */

	err = transaction.Put(idTokenPayload.Sub)
	if err != nil {
		fmt.Println("Error putting transaction:", err.Error())
		err2 := dynamodb.AddIntData(1, food.ID, "food-stock")
		if err2 != nil {
			fmt.Println("Error adding food stock:", err2.Error())
			return response.StatusCode500(err2), nil
		}
		return response.StatusCode500(err), nil
	}

	resBody, err := json.Marshal(transaction)
	if err != nil {
		fmt.Println("Error marshalling transaction:", err.Error())
		return response.StatusCode500(err), nil
	}

	return response.StatusCode200(string(resBody)), nil
}

func main() {
	lambda.Start(handler)
}
