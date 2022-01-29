package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"encoding/json"
	"ubic-food/tools/dynamodb"
	"ubic-food/tools/response"
	"ubic-food/tools/types"
)

type postRequestBody struct {
	Name     string `json:"name"`
	Maker    string `json:"maker"`
	ImageUrl string `json:"imageUrl"`
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	var reqBody postRequestBody
	err := json.Unmarshal([]byte(request.Body), &reqBody)
	if err != nil {
		return response.StatusCode400(err), nil
	}

	dynamodb.CreateTable()
	var food types.Food
	food.ID, err = dynamodb.GenerateID()
	if err != nil {
		return response.StatusCode500(err), nil
	}
	food.Name = reqBody.Name
	food.Maker = reqBody.Maker
	food.ImageUrl = reqBody.ImageUrl
	food.Stock = 0

	err = food.Put()
	if err != nil {
		return response.StatusCode500(err), nil
	}

	resBody, err := json.Marshal(food)
	if err != nil {
		return response.StatusCode500(err), nil
	}

	return response.StatusCode200(string(resBody)), nil
}

func main() {
	lambda.Start(handler)
}
