package callback

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"ubic-food/cookie"
	"ubic-food/hash"

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
	form.Add("redirect_uri", os.Getenv("REDIRECT_URI"))
	form.Add("client_id", os.Getenv("CHANNEL_ID"))
	form.Add("client_secret", os.Getenv("CHANNEL_SECRET"))
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

	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return tokenResponse{}, err
	}

	if resp.StatusCode != 200 {
		return tokenResponse{}, errors.New("Error status code " + strconv.Itoa(resp.StatusCode) + ": " + string(resBody))
	}

	var resBodyStruct tokenResponse
	err = json.Unmarshal(resBody, &resBodyStruct)
	if err != nil {
		return tokenResponse{}, err
	}
	return resBodyStruct, nil
}
