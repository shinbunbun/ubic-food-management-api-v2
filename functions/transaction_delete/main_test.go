package main

import (
	"testing"
	"ubic-food/tools/dynamodb"
	"ubic-food/tools/types"

	"github.com/aws/aws-lambda-go/events"
)

func Test_handler(t *testing.T) {
	res, err := handler(events.APIGatewayProxyRequest{
		Headers: map[string]string{
			"Authorization": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL2FjY2Vzcy5saW5lLm1lIiwic3ViIjoidXNlci0xIiwiYXVkIjoiMTIzNCIsImV4cCI6NDEwMjQ1NTYwMCwiaWF0IjoxNjQzNDUzNzUzLCJub25jZSI6ImR1bW15LW5vbmNlIiwiYW1yIjpbImxpbmVzc28iXSwibmFtZSI6InVzZXItbmFtZSIsInBpY3R1cmUiOiJodHRwczovL2V4YW1wbGUuY29tIn0.HJbFD73yWLe6rXipgynoTf_uV02Z_e2YJxXfWmAlUks",
		},
		PathParameters: map[string]string{
			"transactionId": "transaction-1",
		},
		QueryStringParameters: map[string]string{
			"is_stock_increment": "true",
		},
	})
	if res.StatusCode != 204 {
		t.Fatal("Expected status code 204, got ", res.StatusCode, res.Body)
	}
	if err != nil {
		t.Fatal("Expected no error, got ", err.Error())
	}

	err = dynamodb.AddIntData(-1, "food-1", "food-stock")
	if err != nil {
		t.Fatal("dynamodb.AddIntData error", err.Error())
	}

	var transaction types.Transaction
	transaction.ID = "transaction-1"
	transaction.Date = 1640523361289

	var food types.Food
	food.ID = "food-1"
	food.ImageUrl = "https://shinbunbun.info/images/photos/24.jpeg"
	food.Maker = "dummy-maker-1"
	food.Name = "dummy-name-1"
	food.Stock = 3

	transaction.Food = food
	err = transaction.Put("user-1")
	if err != nil {
		t.Fatal("transaction put failed", err.Error())
	}
}
