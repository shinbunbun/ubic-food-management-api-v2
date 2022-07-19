package main

import (
	"reflect"
	"testing"
	"ubic-food/tools/token"
)

func Test_verify(t *testing.T) {
	type args struct {
		authZHeader string
	}
	tests := []struct {
		name    string
		args    args
		want    token.Payload
		wantErr bool
	}{
		{
			name: "LINE Token",
			args: args{
				authZHeader: "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL2FjY2Vzcy5saW5lLm1lIiwic3ViIjoidXNlci0xIiwiYXVkIjoiMTIzNCIsImV4cCI6NDEwMjQ1NTYwMCwiaWF0IjoxNjQzNDUzNzUzLCJub25jZSI6ImR1bW15LW5vbmNlIiwiYW1yIjpbImxpbmVzc28iXSwibmFtZSI6InVzZXItbmFtZSIsInBpY3R1cmUiOiJodHRwczovL2V4YW1wbGUuY29tIn0.HJbFD73yWLe6rXipgynoTf_uV02Z_e2YJxXfWmAlUks",
			},
			want: token.Payload{
				Iss:   "https://access.line.me",
				Sub:   "U6b53f4ad79a23f5427119cb44f08dbd7",
				Aud:   "1234",
				Exp:   4102455600,
				Iat:   1643453753,
				Nonce: "dummy-nonce",
				Amr: []string{
					"linesso",
				},
				Name:    "user-name",
				Picture: "https://example.com",
			},
			wantErr: false,
		},
		{
			name: "App Token",
			args: args{
				authZHeader: "Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJjbGllbnQtMSIsInN1YiI6InVzZXItMSIsImF1ZCI6InViaWMtZm9vZC5zaGluYnVuYnVuLmluZm8iLCJleHAiOjQxMDI0NTU2MDAsImlhdCI6MTY0MzQ1Mzc1M30.cc2aVj0NegExKtr5M1GmFH0zGAXk9X1vdbRj4ZiA_3P_QYWz-jFRW5h3faZQzL-NKVitHp5yjW_fHhXeFfctyIrGKMFIGaqUCwqvRdtRYGfx_qDMzSWaswqmZAHZSH3tfX9kwl3Fu3KD8_yikSy3K_lzc_rRivNKYfTl6AnEFmOBcpl1wWbMMgKtuHE1sxiH6xBITKpIGlbgxBDViwZ81PWPy45hQBUjgn5hru1J0cSuy8bsujPdLDhyPTHYrDO0Z1EBu2H59rWAr_5iP99V_pJJikiGq5iKp5POVPwja0NrDi3dk7zHfcAstMEu5R_rd5VTlEC-FV8gZovRhHYV_w",
			},
			want: token.Payload{
				Iss: "client-1",
				Sub: "user-1",
				Aud: "ubic-food.shinbunbun.info",
				Exp: 4102455600,
				Iat: 1643453753,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := verify(tt.args.authZHeader)
			if (err != nil) != tt.wantErr {
				t.Errorf("verify() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("verify() = %v, want %v", got, tt.want)
			}
		})
	}
}
