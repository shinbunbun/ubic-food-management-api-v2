package main

/* func Test_handler(t *testing.T) {
	res, err := handler(events.APIGatewayProxyRequest{
		Headers: map[string]string{
			"Authorization": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL2FjY2Vzcy5saW5lLm1lIiwic3ViIjoidXNlci0xIiwiYXVkIjoiMTIzNCIsImV4cCI6NDEwMjQ1NTYwMCwiaWF0IjoxNjQzNDUzNzUzLCJub25jZSI6ImR1bW15LW5vbmNlIiwiYW1yIjpbImxpbmVzc28iXSwibmFtZSI6InVzZXItbmFtZSIsInBpY3R1cmUiOiJodHRwczovL2V4YW1wbGUuY29tIn0.HJbFD73yWLe6rXipgynoTf_uV02Z_e2YJxXfWmAlUks",
		},
	})
	if res.StatusCode != 200 {
		t.Fatal("Expected status code 200, got ", res.StatusCode, res.Body)
	}
	if err != nil {
		t.Fatal("Expected no error, got ", err.Error())
	}

	fmt.Println(res.Body)

	var user types.User
	err = json.Unmarshal([]byte(res.Body), &user)
	if err != nil {
		t.Fatal("response body invalid", err.Error())
	}
	ok := (user.UserID == "user-1") && (user.Name == "user-name")
	if !ok {
		t.Fatal("No expected user data")
	}

	flag := false
	for _, transaction := range user.Transactions {
		if transaction.ID == "transaction-1" {
			ok = (transaction.Date == 1640523361289) && (transaction.Food.ID == "food-1") && (transaction.Food.ImageUrl == "https://shinbunbun.info/images/photos/24.jpeg") && (transaction.Food.Maker == "dummy-maker-1") && (transaction.Food.Name == "dummy-name-1") && (transaction.Food.Stock == 3)
			if !ok {
				t.Fatal("No expected transaction data")
			}
			flag = true
		}
	}
	if !flag {
		t.Fatal("Transaction data is not found")
	}
} */
