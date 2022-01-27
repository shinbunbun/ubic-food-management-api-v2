package food

import (
	"encoding/json"
	"hello-world/dynamodb"
	"hello-world/response"
	"hello-world/token"
	"hello-world/types"

	"github.com/aws/aws-lambda-go/events"
)

type postRequestBody struct {
	Name     string `json:"name"`
	Maker    string `json:"maker"`
	ImageUrl string `json:"imageUrl"`
}

func FoodPost(request events.APIGatewayProxyRequest, idTokenPayload token.Payload) events.APIGatewayProxyResponse {
	var reqBody postRequestBody
	err := json.Unmarshal([]byte(request.Body), &reqBody)
	if err != nil {
		return response.StatusCode400(err)
	}

	dynamodb.CreateTable()
	var food types.Food
	food.ID, err = dynamodb.GenerateID()
	if err != nil {
		return response.StatusCode500(err)
	}
	food.Name = reqBody.Name
	food.Maker = reqBody.Maker
	food.ImageUrl = reqBody.ImageUrl
	food.Stock = 0

	err = food.Put()
	if err != nil {
		return response.StatusCode500(err)
	}

	resBody, err := json.Marshal(food)
	if err != nil {
		return response.StatusCode500(err)
	}

	return response.StatusCode200(string(resBody))

}
