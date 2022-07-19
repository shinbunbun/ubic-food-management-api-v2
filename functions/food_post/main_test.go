package main

/* func Test_handler(t *testing.T) {
	res, err := handler(events.APIGatewayProxyRequest{
		Headers: map[string]string{
			"Authorization": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL2FjY2Vzcy5saW5lLm1lIiwic3ViIjoidXNlci0xIiwiYXVkIjoiMTIzNCIsImV4cCI6NDEwMjQ1NTYwMCwiaWF0IjoxNjQzNDUzNzUzLCJub25jZSI6ImR1bW15LW5vbmNlIiwiYW1yIjpbImxpbmVzc28iXSwibmFtZSI6InVzZXItbmFtZSIsInBpY3R1cmUiOiJodHRwczovL2V4YW1wbGUuY29tIn0.HJbFD73yWLe6rXipgynoTf_uV02Z_e2YJxXfWmAlUks",
		},
		Body: `{"name": "food-2", "maker": "dummy-maker-2", "imageUrl": "https://shinbunbun.info/images/photos/24.jpeg"}`,
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
	ok := (food.ImageUrl == "https://shinbunbun.info/images/photos/24.jpeg") && (food.Maker == "dummy-maker-2") && (food.Name == "food-2") && (food.Stock == 0)
	if !ok {
		t.Fatal("No expected food data")
	}

	var foodVerify types.Food
	foodVerify.ID = food.ID
	err = foodVerify.Get()
	if err != nil {
		t.Fatal("food get failed", err.Error())
	}
	ok = (foodVerify.ImageUrl == "https://shinbunbun.info/images/photos/24.jpeg") && (foodVerify.Maker == "dummy-maker-2") && (foodVerify.Name == "food-2") && (foodVerify.Stock == 0)
	if !ok {
		t.Fatal("No expected food data")
	}
} */
