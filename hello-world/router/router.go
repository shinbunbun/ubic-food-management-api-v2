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
	path := request.Path
	method := request.HTTPMethod
	response := events.APIGatewayProxyResponse{
		StatusCode: 500,
		Body:       "No response",
	}

	var idTokenPayload token.Payload
	var err error
	if !(path == "/auth" || path == "/callback") {
		idTokenPayload, err = authorizer(request)
		if err != nil {
			return events.APIGatewayProxyResponse{
				StatusCode: 401,
				Body:       err.Error(),
			}, nil
		}
	}

	switch path {
	case "/user":
		switch method {
		case "GET":
			response = user.UserGet(request, idTokenPayload)
		}
	case "/transaction":
		switch method {
		case "DELETE":
			/* response =  */ transaction.TransactionDelete()
		case "POST":
			transaction.TransactionPost()
		}
	case "/foods":
		switch method {
		case "GET":
			foods.FoodsGet()
		}
	case "/food":
		switch method {
		case "POST":
			food.FoodPost()
		case "PATCH":
			food.FoodPatch()
		}
	case "/image":
		switch method {
		case "POST":
			image.ImagePost()
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
