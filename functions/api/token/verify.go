package token

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"os"
	"strings"
	"time"
	"ubic-food/functions/api/cookie"
	"ubic-food/functions/api/hash"

	"github.com/aws/aws-lambda-go/events"
)

func VerifyIdToken(requestCookie []string, idToken string) (Payload, error) {
	idTokenArr := strings.Split(idToken, ".")

	err := VerifySignature(idTokenArr)
	if err != nil {
		return Payload{}, err
	}

	idTokenPayload, err := GetIdTokenPayload(idTokenArr)
	if err != nil {
		return Payload{}, err
	}

	err = VerifyIssuer(idTokenPayload)
	if err != nil {
		return Payload{}, err
	}

	err = VerifyAud(idTokenPayload)
	if err != nil {
		return Payload{}, err
	}

	err = VerifyExp(idTokenPayload)
	if err != nil {
		return Payload{}, err
	}

	err = VerifyNonce(requestCookie, idTokenPayload)
	if err != nil {
		return Payload{}, err
	}

	return idTokenPayload, nil
}

func VerifySignature(idTokenArr []string) error {
	validSignatureTarget := idTokenArr[0] + "." + idTokenArr[1]
	signature := idTokenArr[2]
	hmac := base64.RawURLEncoding.EncodeToString(hash.CreateSha256HMAC(validSignatureTarget))
	if signature != hmac {
		return errors.New("Signature is not valid")
	}
	return nil
}

func GetIdTokenPayload(idTokenArr []string) (Payload, error) {
	idTokenPayloadJson, err := base64.StdEncoding.DecodeString(idTokenArr[1])
	if err != nil {
		return Payload{}, err
	}
	var idTokenPayload Payload
	err = json.Unmarshal(idTokenPayloadJson, &idTokenPayload)
	if err != nil {
		return Payload{}, err
	}
	return idTokenPayload, nil
}

func GetIdTokenPayloadByAuthZHeader(request events.APIGatewayProxyRequest) (Payload, error) {
	authZHeader := request.Headers["Authorization"]
	idToken := strings.Split(authZHeader, "Bearer ")[1]
	idTokenArr := strings.Split(idToken, ".")
	idTokenPayload, err := GetIdTokenPayload(idTokenArr)
	if err != nil {
		return Payload{}, err
	}
	return idTokenPayload, nil
}

func VerifyIssuer(idTokenPayload Payload) error {
	issuer := "https://access.line.me"
	if idTokenPayload.Iss != issuer {
		return errors.New("Issuer is not valid")
	}
	return nil
}

func VerifyAud(idTokenPayload Payload) error {
	if idTokenPayload.Aud != os.Getenv("CHANNEL_ID") {
		return errors.New("Aud is not valid")
	}
	return nil
}

func VerifyExp(idTokenPayload Payload) error {
	if idTokenPayload.Exp < int(time.Now().Unix()) {
		return errors.New("Token is expired")
	}
	return nil
}

func VerifyNonce(requestCookie []string, idTokenPayload Payload) error {
	cookieNonce, err := cookie.GetCookieValue(requestCookie, "nonce")
	if err != nil {
		return err
	}
	cookieNonceHash := hash.CreateSha3_256Hash(cookieNonce)
	if cookieNonceHash != idTokenPayload.Nonce {
		return errors.New("State is not valid")
	}
	return nil
}
