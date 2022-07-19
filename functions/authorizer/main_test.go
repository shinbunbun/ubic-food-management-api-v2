package main

import (
	"context"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func Test_handler(t *testing.T) {
	requests := []events.APIGatewayCustomAuthorizerRequest{
		{
			AuthorizationToken: "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL2FjY2Vzcy5saW5lLm1lIiwic3ViIjoidXNlci0xIiwiYXVkIjoiMTIzNCIsImV4cCI6NDEwMjQ1NTYwMCwiaWF0IjoxNjQzNDUzNzUzLCJub25jZSI6ImR1bW15LW5vbmNlIiwiYW1yIjpbImxpbmVzc28iXSwibmFtZSI6InVzZXItbmFtZSIsInBpY3R1cmUiOiJodHRwczovL2V4YW1wbGUuY29tIn0.HJbFD73yWLe6rXipgynoTf_uV02Z_e2YJxXfWmAlUks",
			MethodArn:          "arn:aws:test-arn",
		},
		{
			AuthorizationToken: "Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJjbGllbnQtMSIsInN1YiI6InVzZXItMSIsImF1ZCI6InViaWMtZm9vZC5zaGluYnVuYnVuLmluZm8iLCJleHAiOjQxMDI0NTU2MDAsImlhdCI6MTY0MzQ1Mzc1M30.cc2aVj0NegExKtr5M1GmFH0zGAXk9X1vdbRj4ZiA_3P_QYWz-jFRW5h3faZQzL-NKVitHp5yjW_fHhXeFfctyIrGKMFIGaqUCwqvRdtRYGfx_qDMzSWaswqmZAHZSH3tfX9kwl3Fu3KD8_yikSy3K_lzc_rRivNKYfTl6AnEFmOBcpl1wWbMMgKtuHE1sxiH6xBITKpIGlbgxBDViwZ81PWPy45hQBUjgn5hru1J0cSuy8bsujPdLDhyPTHYrDO0Z1EBu2H59rWAr_5iP99V_pJJikiGq5iKp5POVPwja0NrDi3dk7zHfcAstMEu5R_rd5VTlEC-FV8gZovRhHYV_w",
			MethodArn:          "arn:aws:test-arn",
		},
	}

	for _, request := range requests {
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
}
