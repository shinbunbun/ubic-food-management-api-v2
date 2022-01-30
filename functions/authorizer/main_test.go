package main

import (
	"context"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func Test_handler(t *testing.T) {
	request := events.APIGatewayCustomAuthorizerRequest{
		AuthorizationToken: "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL2FjY2Vzcy5saW5lLm1lIiwic3ViIjoidXNlci0xIiwiYXVkIjoiMTIzNCIsImV4cCI6NDEwMjQ1NTYwMCwiaWF0IjoxNjQzNDUzNzUzLCJub25jZSI6ImR1bW15LW5vbmNlIiwiYW1yIjpbImxpbmVzc28iXSwibmFtZSI6InVzZXItbmFtZSIsInBpY3R1cmUiOiJodHRwczovL2V4YW1wbGUuY29tIn0.HJbFD73yWLe6rXipgynoTf_uV02Z_e2YJxXfWmAlUks",
		MethodArn:          "arn:aws:test-arn",
	}

	var ctx context.Context
	response, err := handler(ctx, request)
	if err != nil {
		t.Errorf("handler returned error: %v", err)
	}
	ok := (response.PrincipalID == "user-1") && (response.PolicyDocument.Statement[0].Action[0] == "execute-api:Invoke") && (response.PolicyDocument.Statement[0].Effect == "Allow") && (response.PolicyDocument.Statement[0].Resource[0] == "arn:aws:test-arn")
	if !ok {
		t.Errorf("handler returned invalid response: %v", response)
	}
}
