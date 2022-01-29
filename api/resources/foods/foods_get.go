package foods

import (
	"encoding/json"
	"fmt"
	"ubic-food/api/dynamodb"
	"ubic-food/api/response"
	"ubic-food/api/token"
	"ubic-food/api/types"

	"github.com/aws/aws-lambda-go/events"
)

func FoodsGet(request events.APIGatewayProxyRequest, idTokenPayload token.Payload) events.APIGatewayProxyResponse {

	dynamodb.CreateTable()

	items, err := dynamodb.GetByDataKind("food")
	if err != nil {
		return response.StatusCode500(err)
	}

	fmt.Printf("item: %+v\n", items)

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
		return response.StatusCode500(err)
	}

	return response.StatusCode200(string(resBody))
}
