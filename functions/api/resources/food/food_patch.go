package food

import (
	"encoding/json"
	"ubic-food/functions/api/dynamodb"
	"ubic-food/functions/api/response"
	"ubic-food/functions/api/token"
	"ubic-food/functions/api/types"

	"github.com/aws/aws-lambda-go/events"
)

type patchRequestBody struct {
	AddNum int `json:"addNum"`
}

func FoodPatch(request events.APIGatewayProxyRequest, idTokenPayload token.Payload) events.APIGatewayProxyResponse {
	reqBodyJson := request.Body
	var reqBody patchRequestBody
	err := json.Unmarshal([]byte(reqBodyJson), &reqBody)
	if err != nil {
		return response.StatusCode400(err)
	}
	addNum := reqBody.AddNum

	foodId := request.PathParameters["foodId"]

	dynamodb.CreateTable()
	err = dynamodb.AddIntData(addNum, foodId, "food-stock")
	if err != nil {
		return response.StatusCode500(err)
	}

	var food types.Food
	food.ID = foodId
	err = food.Get()
	if err != nil {
		return response.StatusCode500(err)
	}

	resBody, err := json.Marshal(food)
	if err != nil {
		return response.StatusCode500(err)
	}

	return response.StatusCode200(string(resBody))
}
