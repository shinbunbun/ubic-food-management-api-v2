package main

import (
	"encoding/json"
	"ubic-food/tools/dynamodb"
	"ubic-food/tools/response"
	"ubic-food/tools/types"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type patchRequestBody struct {
	AddNum int `json:"addNum"`
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	reqBodyJson := request.Body
	var reqBody patchRequestBody
	err := json.Unmarshal([]byte(reqBodyJson), &reqBody)
	if err != nil {
		return response.StatusCode400(err), nil
	}
	addNum := reqBody.AddNum

	foodId := request.PathParameters["foodId"]

	err = dynamodb.AddIntData(addNum, foodId, "food-stock")
	if err != nil {
		return response.StatusCode500(err), nil
	}

	var food types.Food
	food.ID = foodId
	err = food.Get()
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
