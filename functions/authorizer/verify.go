package main

import (
	"strings"
	"ubic-food/functions/api/token"
)

func verify(authZHeader string) (token.Payload, error) {
	idToken := strings.Split(authZHeader, "Bearer ")[1]
	idTokenArr := strings.Split(idToken, ".")
	err := token.VerifySignature(idTokenArr)
	if err != nil {
		return token.Payload{}, err
	}

	idTokenPayload, err := token.GetIdTokenPayload(idTokenArr)
	if err != nil {
		return token.Payload{}, err
	}

	err = token.VerifyIssuer(idTokenPayload)
	if err != nil {
		return token.Payload{}, err
	}

	err = token.VerifyAud(idTokenPayload)
	if err != nil {
		return token.Payload{}, err
	}

	err = token.VerifyExp(idTokenPayload)
	if err != nil {
		return token.Payload{}, err
	}

	return idTokenPayload, nil
}
