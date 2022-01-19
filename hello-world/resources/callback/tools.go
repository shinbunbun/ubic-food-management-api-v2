package callback

import (
	"encoding/json"
	"errors"
	"hello-world/config"
	"hello-world/cookie"
	"hello-world/hash"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/aws/aws-lambda-go/events"
)

func checkState(query map[string]string, requestCookie []string, request events.APIGatewayProxyRequest) error {
	callbackStateHash := query["state"]
	cookieState, err := cookie.GetCookieValue(requestCookie, "state")
	if err != nil {
		return err
	}
	cookieStateHash := hash.CreateSha3_256Hash(cookieState)
	if cookieStateHash != callbackStateHash {
		return errors.New("State is not valid")
	}
	return nil
}

func getAccessToken(code string) (tokenResponse, error) {

	form := url.Values{}
	form.Add("grant_type", "authorization_code")
	form.Add("code", code)
	form.Add("redirect_uri", config.GetRedirectUri())
	form.Add("client_id", config.GetEnv("CHANNEL_ID"))
	form.Add("client_secret", config.GetEnv("CHANNEL_SECRET"))
	body := strings.NewReader(form.Encode())

	reqUrl := "https://api.line.me/oauth2/v2.1/token"
	req, err := http.NewRequest("POST", reqUrl, body)
	if err != nil {
		return tokenResponse{}, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return tokenResponse{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return tokenResponse{}, errors.New("Error status code " + string(resp.StatusCode) + ": " + resp.Status)
	}

	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return tokenResponse{}, err
	}

	var resBodyStruct tokenResponse
	json.Unmarshal(resBody, &resBodyStruct)
	return resBodyStruct, nil
}
