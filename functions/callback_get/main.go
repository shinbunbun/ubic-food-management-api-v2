package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"ubic-food/functions/api/cookie"
	"ubic-food/functions/api/hash"
	"ubic-food/functions/api/response"
	"ubic-food/functions/api/token"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type tokenResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
	Scope        string `json:"scope"`
	IdToken      string `json:"id_token"`
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	query := request.QueryStringParameters

	headers := http.Header{}
	for header, values := range request.MultiValueHeaders {
		for _, value := range values {
			headers.Add(header, value)
		}
	}

	requestCookie := strings.Split(headers.Get("cookie"), "; ")

	err := checkState(query, requestCookie, request)
	if err != nil {
		fmt.Println("State Error:", err)
		return response.StatusCode400(err), nil
	}

	code := query["code"]

	tokenRes, err := getAccessToken(code)
	if err != nil {
		fmt.Println("Get Token Error:", err)
		return response.StatusCode500(err), nil
	}
	idToken := tokenRes.IdToken

	_, err = token.VerifyIdToken(requestCookie, idToken)
	if err != nil {
		fmt.Println("Verify Token Error:", err)
		return response.StatusCode500(err), nil
	}

	return response.StatusCode200(idToken), nil
}

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

func main() {
	lambda.Start(handler)
}
