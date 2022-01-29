package image

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"ubic-food/functions/api/dynamodb"
	"ubic-food/functions/api/response"
	"ubic-food/functions/api/token"

	"ubic-food/functions/api/s3"

	"github.com/aws/aws-lambda-go/events"
)

type imagePostResponse struct {
	ImageUrl string `json:"imageUrl"`
}

func ImagePost(request events.APIGatewayProxyRequest, idTokenPayload token.Payload) events.APIGatewayProxyResponse {
	reqBody := request.Body
	dec, err := base64.StdEncoding.DecodeString(reqBody)
	if err != nil {
		return response.StatusCode500(err)
	}

	fileName, err := dynamodb.GenerateID()
	if err != nil {
		return response.StatusCode500(err)
	}
	fileName += ".jpeg"

	var buf bytes.Buffer
	buf.Write(dec)

	location, err := s3.Upload(&buf, fileName, request.Headers["Content-Type"])
	if err != nil {
		return response.StatusCode500(err)
	}

	res, err := json.Marshal(imagePostResponse{
		ImageUrl: location,
	})
	if err != nil {
		return response.StatusCode500(err)
	}

	return response.StatusCode200(string(res))

}
