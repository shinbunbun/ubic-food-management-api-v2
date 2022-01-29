package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"ubic-food/functions/api/dynamodb"
	"ubic-food/functions/api/response"

	"ubic-food/functions/api/s3"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type imagePostResponse struct {
	ImageUrl string `json:"imageUrl"`
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	reqBody := request.Body
	dec, err := base64.StdEncoding.DecodeString(reqBody)
	if err != nil {
		return response.StatusCode500(err), nil
	}

	fileName, err := dynamodb.GenerateID()
	if err != nil {
		return response.StatusCode500(err), nil
	}
	fileName += ".jpeg"

	var buf bytes.Buffer
	buf.Write(dec)

	location, err := s3.Upload(&buf, fileName, request.Headers["Content-Type"])
	if err != nil {
		return response.StatusCode500(err), nil
	}

	res, err := json.Marshal(imagePostResponse{
		ImageUrl: location,
	})
	if err != nil {
		return response.StatusCode500(err), nil
	}

	return response.StatusCode200(string(res)), nil
}

func main() {
	lambda.Start(handler)
}
