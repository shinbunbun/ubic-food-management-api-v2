package main

import (
	"encoding/json"
	"strings"
	"ubic-food/tools/keypair"
	"ubic-food/tools/token"
)

func verify(authZHeader string) (token.Payload, error) {
	idToken := strings.Split(authZHeader, "Bearer ")[1]

	keyPair := keypair.KeyPair{}
	claims, err := keyPair.Verify(idToken)
	if err == nil {
		claimJson, err := json.Marshal(claims)
		if err != nil {
			return token.Payload{}, err
		}

		var payload token.Payload
		err = json.Unmarshal(claimJson, &payload)
		if err != nil {
			return token.Payload{}, err
		}

		return payload, nil
	} /*  else {
		println()
	} */

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
