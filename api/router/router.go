package router

import (
	"strings"
	"ubic-food/api/resources/auth"
	"ubic-food/api/resources/callback"
	"ubic-food/api/resources/food"
	"ubic-food/api/resources/foods"
	"ubic-food/api/resources/image"
	"ubic-food/api/resources/transaction"
	"ubic-food/api/resources/user"
	"ubic-food/api/response"
	"ubic-food/api/token"

	"github.com/aws/aws-lambda-go/events"
)

func Router(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	resource := request.Resource
	method := request.HTTPMethod
	var res events.APIGatewayProxyResponse

	var idTokenPayload token.Payload
	var err error
	if !(resource == "/auth" || resource == "/callback") {
		authZHeader := request.Headers["Authorization"]
		idToken := strings.Split(authZHeader, "Bearer ")[1]
		idTokenArr := strings.Split(idToken, ".")
		idTokenPayload, err = token.GetIdTokenPayload(idTokenArr)
		if err != nil {
			return response.StatusCode500(err), err
		}
	}

	switch resource {
	case "/user":
		switch method {
		case "GET":
			res = user.UserGet(request, idTokenPayload)
		}
	case "/transaction/{transactionId}":
		switch method {
		case "DELETE":
			res = transaction.TransactionDelete(request, idTokenPayload)
		}
	case "/transaction":
		switch method {
		case "POST":
			res = transaction.TransactionPost(request, idTokenPayload)
		}
	case "/foods":
		switch method {
		case "GET":
			res = foods.FoodsGet(request, idTokenPayload)
		}
	case "/food":
		switch method {
		case "POST":
			res = food.FoodPost(request, idTokenPayload)
		}
	case "/food/{foodId}":
		switch method {
		case "PATCH":
			res = food.FoodPatch(request, idTokenPayload)
		}
	case "/image":
		switch method {
		case "POST":
			res = image.ImagePost(request, idTokenPayload)
		}
	case "/auth":
		switch method {
		case "GET":
			res = auth.AuthGet(request)
		}
	case "/callback":
		switch method {
		case "GET":
			res = callback.CallbackGet(request)
		}
	}
	return res, nil
}
