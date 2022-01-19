package callback

type tokenResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
	Scope        string `json:"scope"`
	IdToken      string `json:"id_token"`
}

type payload struct {
	Iss      string   `json:"iss"`
	Sub      string   `json:"sub"`
	Aud      string   `json:"aud"`
	Exp      int      `json:"exp"`
	Iat      int      `json:"iat"`
	AuthTime int      `json:"auth_time"`
	Nonce    string   `json:"nonce"`
	Amr      []string `json:"amr,omitempty"`
	Name     string   `json:"name"`
	Picture  string   `json:"picture"`
	Email    string   `json:"email,omitempty"`
}
