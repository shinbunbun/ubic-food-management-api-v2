package token

type Payload struct {
	Iss string `json:"iss"`
	Sub string `json:"sub"`
	Aud string `json:"aud"`
	Exp int    `json:"exp"`
	Iat int    `json:"iat"`
	// AuthTime int      `json:"auth_time"`
	Nonce   string   `json:"nonce,omitempty"`
	Amr     []string `json:"amr,omitempty"`
	Name    string   `json:"name,omitempty"`
	Picture string   `json:"picture,omitempty"`
	Email   string   `json:"email,omitempty"`
}
