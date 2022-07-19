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
		Headers: map[string]string{
			"Authorization": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL2FjY2Vzcy5saW5lLm1lIiwic3ViIjoiVTZiNTNmNGFkNzlhMjNmNTQyNzExOWNiNDRmMDhkYmQ3IiwiYXVkIjoiMTIzNCIsImV4cCI6NDEwMjQ1NTYwMCwiaWF0IjoxNjQzNDUzNzUzLCJub25jZSI6ImR1bW15LW5vbmNlIiwiYW1yIjpbImxpbmVzc28iXSwibmFtZSI6InVzZXItbmFtZSIsInBpY3R1cmUiOiJodHRwczovL2V4YW1wbGUuY29tIn0.VRNCuRKWgmAazYopeDDbL1PINOt58cXy2HoyvHfXmfo",
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
