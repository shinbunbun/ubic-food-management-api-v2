package token

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"hello-world/config"
	"hello-world/cookie"
	"hello-world/hash"
	"strings"
	"time"
)

func VerifyIdToken(requestCookie []string, idToken string) (payload, error) {
	idTokenArr := strings.Split(idToken, ".")

	err := verifySignature(idTokenArr)
	if err != nil {
		return payload{}, err
	}

	idTokenPayload, err := getIdTokenPayload(idTokenArr)
	if err != nil {
		return payload{}, err
	}

	err = verifyIssuer(idTokenPayload)
	if err != nil {
		return payload{}, err
	}

	err = verifyAud(idTokenPayload)
	if err != nil {
		return payload{}, err
	}

	err = verifyExp(idTokenPayload)
	if err != nil {
		return payload{}, err
	}

	err = verifyNonce(requestCookie, idTokenPayload)
	if err != nil {
		return payload{}, err
	}

	return idTokenPayload, nil
}

func verifySignature(idTokenArr []string) error {
	validSignatureTarget := idTokenArr[0] + "." + idTokenArr[1]
	signature := idTokenArr[2]
	hmac := base64.RawURLEncoding.EncodeToString(hash.CreateSha256HMAC(validSignatureTarget))
	if signature != hmac {
		return errors.New("Signature is not valid")
	}
	return nil
}

func getIdTokenPayload(idTokenArr []string) (payload, error) {
	idTokenPayloadJson, err := base64.StdEncoding.DecodeString(idTokenArr[1])
	if err != nil {
		return payload{}, err
	}
	var idTokenPayload payload
	err = json.Unmarshal(idTokenPayloadJson, &idTokenPayload)
	if err != nil {
		return payload{}, err
	}
	return idTokenPayload, nil
}

func verifyIssuer(idTokenPayload payload) error {
	issuer := "https://access.line.me"
	if idTokenPayload.Iss != issuer {
		return errors.New("Issuer is not valid")
	}
	return nil
}

func verifyAud(idTokenPayload payload) error {
	if idTokenPayload.Aud != config.GetEnv("CHANNEL_ID") {
		return errors.New("Aud is not valid")
	}
	return nil
}

func verifyExp(idTokenPayload payload) error {
	if idTokenPayload.Exp < int(time.Now().Unix()) {
		return errors.New("Token is expired")
	}
	return nil
}

func verifyNonce(requestCookie []string, idTokenPayload payload) error {
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
