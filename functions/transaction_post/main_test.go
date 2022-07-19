package main

/* func Test_handler(t *testing.T) {

	request := events.APIGatewayProxyRequest{
		Body: `{"foodId": "food-1"}`,
		QueryStringParameters: map[string]string{
			"is_stock_decrement": "true",
		},
		Headers: map[string]string{
			"Authorization": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL2FjY2Vzcy5saW5lLm1lIiwic3ViIjoidXNlci0xIiwiYXVkIjoiMTIzNCIsImV4cCI6NDEwMjQ1NTYwMCwiaWF0IjoxNjQzNDUzNzUzLCJub25jZSI6ImR1bW15LW5vbmNlIiwiYW1yIjpbImxpbmVzc28iXSwibmFtZSI6InVzZXItbmFtZSIsInBpY3R1cmUiOiJodHRwczovL2V4YW1wbGUuY29tIn0.HJbFD73yWLe6rXipgynoTf_uV02Z_e2YJxXfWmAlUks",
		},
	}

	res, err := handler(request)
	if res.StatusCode != 200 {
		t.Fatal("Expected status code 200, got ", res.StatusCode)
	}
	if err != nil {
		t.Fatal("Expected no error, got ", err.Error())
	}

	fmt.Println(res.Body)

	var transaction types.Transaction
	err = json.Unmarshal([]byte(res.Body), &transaction)
	if err != nil {
		t.Fatal("response body invalid", err.Error())
	}

	err = dynamodb.AddIntData(1, "food-1", "food-stock")
	if err != nil {
		t.Fatal("dynamodb.AddIntData error", err.Error())
	}
	transaction.Food.Stock += 1

	fmt.Printf("transaction: %+#v\n", transaction)

	ok := (transaction.Food.ID == "food-1") && (transaction.Food.ImageUrl == "https://shinbunbun.info/images/photos/24.jpeg") && (transaction.Food.Maker == "dummy-maker-1") && (transaction.Food.Name == "dummy-name-1") && (transaction.Food.Stock == 3)
	if !ok {
		t.Fatal("No expected food data")
	}
} */
