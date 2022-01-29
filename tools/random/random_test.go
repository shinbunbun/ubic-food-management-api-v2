package random

import (
	"fmt"
	"testing"
)

func TestGenerateRandomString(t *testing.T) {
	type args struct {
		s int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success test 1",
			args: args{
				s: 32,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenerateRandomString(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateRandomString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println(got)
		})
	}
}
