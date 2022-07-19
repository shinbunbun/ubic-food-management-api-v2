package main

import (
	"errors"
	"strings"
	"ubic-food/tools/keypair"
	"ubic-food/tools/token"
)

func verify(authZHeader string) (token.Payload, error) {
	idToken := strings.Split(authZHeader, "Bearer ")[1]

	keyPair := keypair.KeyPair{}
	claims, err := keyPair.Verify(idToken)
	if err == nil {
		println("Verified token:", claims)
		payload, ok := claims.(token.Payload)
		if !ok {
			return token.Payload{}, errors.New("Invalid claims")
		}
		return payload, nil
	} else {
		println(err.Error())
	}

	idTokenArr := strings.Split(idToken, ".")
	err = token.VerifySignature(idTokenArr)
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
