package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"ubic-food/tools/types"

	"github.com/aws/aws-lambda-go/events"
)

func Test_handler(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SP0R2USEDHqPV7mcIK08ZAs4WtPMQ0NdMHuSD8tnWOw")
		fmt.Fprintf(w, "127.0.0.1")
	}))
	defer ts.Close()

	res, err := handler(events.APIGatewayProxyRequest{})
	if res.StatusCode != 200 {
		t.Fatal("Expected status code 200, got ", res.StatusCode)
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
