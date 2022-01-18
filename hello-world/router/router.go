package router

import (
	"hello-world/resources/auth"
	"hello-world/resources/food"
	"hello-world/resources/foods"
	"hello-world/resources/image"
	"hello-world/resources/transaction"
	"hello-world/resources/user"

	"github.com/aws/aws-lambda-go/events"
)

func Router(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	path := request.Path
	method := request.HTTPMethod
	var response events.APIGatewayProxyResponse
	switch path {
	case "/user":
		switch method {
		case "GET":
			/* response =  */ user.UserGet()
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
	}
	return response, nil
}
