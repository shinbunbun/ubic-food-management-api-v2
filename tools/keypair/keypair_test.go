package keypair

import (
	"testing"
)

func TestKeyPair_Generate(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "success",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := KeyPair{}
			err := k.Generate()
			if (err != nil) != tt.wantErr {
				t.Errorf("KeyPair.Generate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
