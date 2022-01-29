package router

import (
	"ubic-food/functions/api/resources/callback"

	"github.com/aws/aws-lambda-go/events"
)

func Router(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	resource := request.Resource
	method := request.HTTPMethod
	var res events.APIGatewayProxyResponse

	switch resource {
	case "/callback":
		switch method {
		case "GET":
			res = callback.CallbackGet(request)
		}
	}
	return res, nil
}
