package main

import (
	"encoding/json"
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
			"foodId": "food-1",
		},
		Body: `{"addNum": 1}`,
	})
	if res.StatusCode != 200 {
		t.Fatal("Expected status code 200, got ", res.StatusCode, res.Body)
	}
	if err != nil {
		t.Fatal("Expected no error, got ", err.Error())
	}

	var food types.Food

	err = json.Unmarshal([]byte(res.Body), &food)
	if err != nil {
		t.Fatal("response body invalid", err.Error())
	}
	ok := (food.ID == "food-1") && (food.ImageUrl == "https://shinbunbun.info/images/photos/24.jpeg") && (food.Maker == "dummy-maker-1") && (food.Name == "dummy-name-1") && (food.Stock == 4)
	if !ok {
		t.Fatal("No expected food data")
	}

	var foodVerify types.Food
	foodVerify.ID = food.ID
	err = foodVerify.Get()
	if err != nil {
		t.Fatal("food get failed", err.Error())
	}
	ok = (foodVerify.ImageUrl == "https://shinbunbun.info/images/photos/24.jpeg") && (foodVerify.Maker == "dummy-maker-1") && (foodVerify.Name == "dummy-name-1") && (foodVerify.Stock == 4)
	if !ok {
		t.Fatal("No expected food data")
	}

	err = dynamodb.AddIntData(-1, food.ID, "food-stock")
	if err != nil {
		t.Fatal("food stock add failed", err.Error())
	}
}
