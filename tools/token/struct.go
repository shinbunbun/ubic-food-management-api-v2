package token

type Payload struct {
	Iss string `json:"iss"`
	Sub string `json:"sub"`
	Aud string `json:"aud"`
	Exp int    `json:"exp"`
	Iat int    `json:"iat"`
	// AuthTime int      `json:"auth_time"`
	Nonce   string   `json:"nonce"`
	Amr     []string `json:"amr,omitempty"`
	Name    string   `json:"name"`
	Picture string   `json:"picture"`
	Email   string   `json:"email,omitempty"`
}

func (p Payload) Valid() error {
	return nil
}
