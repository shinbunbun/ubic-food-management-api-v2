package main

import (
	"encoding/json"
	"fmt"
	"testing"
	"ubic-food/tools/types"

	"github.com/aws/aws-lambda-go/events"
)

func Test_handler(t *testing.T) {
	res, err := handler(events.APIGatewayProxyRequest{})
	if res.StatusCode != 200 {
		t.Fatal("Expected status code 200, got ", res.StatusCode, res.Body)
	}
	if err != nil {
		t.Fatal("Expected no error, got ", err.Error())
	}

	fmt.Println(res.Body)

	foods := make([]types.Food, 0)
	err = json.Unmarshal([]byte(res.Body), &foods)
	if err != nil {
		t.Fatal("response body invalid", err.Error())
	}
	flag := false
	for _, food := range foods {
		if food.ID == "food-1" {
			flag = true
			ok := (food.ImageUrl == "https://shinbunbun.info/images/photos/24.jpeg") && (food.Maker == "dummy-maker-1") && (food.Name == "dummy-name-1") && (food.Stock == 3)
			if !ok {
				t.Fatal("No expected food data")
			}
		}
	}
	if !flag {
		t.Fatal("Expected food data is not found")
	}
}
