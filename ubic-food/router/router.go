package router

import (
	"hello-world/resources/auth"
	"hello-world/resources/callback"
	"hello-world/resources/food"
	"hello-world/resources/foods"
	"hello-world/resources/image"
	"hello-world/resources/transaction"
	"hello-world/resources/user"
	"hello-world/token"

	"github.com/aws/aws-lambda-go/events"
)

func Router(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	resource := request.Resource
	method := request.HTTPMethod
	response := events.APIGatewayProxyResponse{
		StatusCode: 500,
		Body:       "No response",
	}

	var idTokenPayload token.Payload
	var err error
	if !(resource == "/auth" || resource == "/callback") {
		idTokenPayload, err = authorizer(request)
		if err != nil {
			return events.APIGatewayProxyResponse{
				StatusCode: 401,
				Body:       err.Error(),
			}, nil
		}
	}

	switch resource {
	case "/user":
		switch method {
		case "GET":
			response = user.UserGet(request, idTokenPayload)
		}
	case "/transaction/{transactionId}":
		switch method {
		case "DELETE":
			response = transaction.TransactionDelete(request, idTokenPayload)
		}
	case "/transaction":
		switch method {
		case "POST":
			response = transaction.TransactionPost(request, idTokenPayload)
		}
	case "/foods":
		switch method {
		case "GET":
			response = foods.FoodsGet(request, idTokenPayload)
		}
	case "/food":
		switch method {
		case "POST":
			response = food.FoodPost(request, idTokenPayload)
		}
	case "/food/{foodId}":
		switch method {
		case "PATCH":
			response = food.FoodPatch(request, idTokenPayload)
		}
	case "/image":
		switch method {
		case "POST":
			response = image.ImagePost(request, idTokenPayload)
		}
	case "/auth":
		switch method {
		case "GET":
			response = auth.AuthGet(request)
		}
	case "/callback":
		switch method {
		case "GET":
			response = callback.CallbackGet(request)
		}
	}
	return response, nil
}
