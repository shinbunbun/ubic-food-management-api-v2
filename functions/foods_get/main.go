package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"encoding/json"
	"ubic-food/tools/dynamodb"
	"ubic-food/tools/response"
	"ubic-food/tools/types"
)

func handler(_ events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	items, err := dynamodb.GetByDataKind("food")
	if err != nil {
		return response.StatusCode500(err), nil
	}

	var foodsMap = make(map[string]types.Food)
	for _, v := range items {
		_, ok := foodsMap[v.ID]
		if !ok {
			foodsMap[v.ID] = types.Food{
				ID: v.ID,
			}
		}
		food := foodsMap[v.ID]
		if v.DataType == "food-image" {
			food.ImageUrl = v.Data
		}
		if v.DataType == "food-maker" {
			food.Maker = v.Data
		}
		if v.DataType == "food-name" {
			food.Name = v.Data
		}
		if v.DataType == "food-stock" {
			food.Stock = *(v.IntData)
		}
		foodsMap[v.ID] = food
	}

	var foods []types.Food
	for _, v := range foodsMap {
		foods = append(foods, v)
	}

	resBody, err := json.Marshal(foods)
	if err != nil {
		return response.StatusCode500(err), nil
	}

	return response.StatusCode200(string(resBody)), nil
}

func main() {
	lambda.Start(handler)
}
