package main

import (
	"context"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func Test_handler(t *testing.T) {
	tests := []struct {
		name    string
		request events.APIGatewayCustomAuthorizerRequest
	}{
		{
			name: "LINE Token",
			request: events.APIGatewayCustomAuthorizerRequest{
				AuthorizationToken: "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL2FjY2Vzcy5saW5lLm1lIiwic3ViIjoidXNlci0xIiwiYXVkIjoiMTIzNCIsImV4cCI6NDEwMjQ1NTYwMCwiaWF0IjoxNjQzNDUzNzUzLCJub25jZSI6ImR1bW15LW5vbmNlIiwiYW1yIjpbImxpbmVzc28iXSwibmFtZSI6InVzZXItbmFtZSIsInBpY3R1cmUiOiJodHRwczovL2V4YW1wbGUuY29tIn0.HJbFD73yWLe6rXipgynoTf_uV02Z_e2YJxXfWmAlUks",
				MethodArn:          "arn:aws:test-arn",
			},
		},
		{
			name: "Apps Token",
			request: events.APIGatewayCustomAuthorizerRequest{
				AuthorizationToken: "Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImlkIjoiY2xpZW50LTEifQ.eyJpc3MiOiJjbGllbnQtMSIsInN1YiI6InVzZXItMSIsImF1ZCI6InViaWMtZm9vZC5zaGluYnVuYnVuLmluZm8iLCJleHAiOjQxMDI0NTU2MDAsImlhdCI6MTY0MzQ1Mzc1M30.rBLYRrY-fhqu4k4vVeOjB-Lj9Acq-ooR5wm094gh9nnbHmRtw8sChynVYMKQ7uswwz3KICC1nIoztcijGKb3iYFAiRDI0ZEc6qBYt-pqam7--jTYnFTiUgWK1xRljPSFnBLRkWNzptQ1z13Vc4V8nN-CgDNsLzy58If85PsLS25_EuLKS6YybPpw3tIXb00GLTcnFQ5g5UaHkCo5RhgW5peuCSTMpPnEF2og4_abt-zKpY882MO8bzFE1jfJh_YGMFWN59tE8iTfnZwN9ZsSGH4QdAFUBw_n_MQlP6QhP9N9KdVdh56gPKlvPMWRNrWOp6-cyTzmC1i6GDialcgxpQ",
				MethodArn:          "arn:aws:test-arn",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var ctx context.Context
			response, err := handler(ctx, tt.request)
			if err != nil {
				t.Errorf("handler returned error: %v", err)
			}
			ok := (response.PrincipalID == "user-1") && (response.PolicyDocument.Statement[0].Action[0] == "execute-api:Invoke") && (response.PolicyDocument.Statement[0].Effect == "Allow") && (response.PolicyDocument.Statement[0].Resource[0] == "arn:aws:test-arn")
			if !ok {
				t.Errorf("handler returned invalid response: %v", response)
			}
		})
	}
}
