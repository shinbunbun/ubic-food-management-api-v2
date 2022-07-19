package main

import (
	"encoding/json"
	"testing"
	"ubic-food/tools/dynamodb"
	"ubic-food/tools/keypair"

	"github.com/aws/aws-lambda-go/events"
)

func Test_handler(t *testing.T) {
	res, err := handler(events.APIGatewayProxyRequest{
		QueryStringParameters: map[string]string{
			"clientId": "client-1",
		},
	})
	if res.StatusCode != 200 {
		t.Fatal("Expected status code 200, got ", res.StatusCode, res.Body)
	}
	if err != nil {
		t.Fatal("Expected no error, got ", err.Error())
	}

	var keyPair keypair.KeyPair
	err = json.Unmarshal([]byte(res.Body), &keyPair)
	if err != nil {
		t.Fatal("response body invalid", err.Error())
	}

	keyData, err := dynamodb.GetByIDDataType("client-1", "public-key")
	if err != nil {
		t.Fatal("Expected no error, got ", err.Error())
	}
	if keyData.ID != "client-1" {
		t.Fatal("Unexpected ID, got ", keyData.ID)
	}

}
