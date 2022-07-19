package keypair

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"time"
	"ubic-food/tools/dynamodb"

	"github.com/dgrijalva/jwt-go"
)

type KeyPair struct {
	PublicKey  string
	PrivateKey string
}

func (k *KeyPair) Generate() error {
	reader := rand.Reader
	bitSize := 2048

	rsaPrivateKey, err := rsa.GenerateKey(reader, bitSize)
	if err != nil {
		return err
	}

	derRsaPrivateKey := x509.MarshalPKCS1PrivateKey(rsaPrivateKey)
	var buf1 bytes.Buffer
	err = pem.Encode(&buf1, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: derRsaPrivateKey})
	if err != nil {
		return err
	}
	k.PrivateKey = buf1.String()

	serRsaPublicKey := x509.MarshalPKCS1PublicKey(&rsaPrivateKey.PublicKey)
	var buf2 bytes.Buffer
	err = pem.Encode(&buf2, &pem.Block{Type: "RSA PUBLIC KEY", Bytes: serRsaPublicKey})
	k.PublicKey = buf2.String()

	return err
}

func (k *KeyPair) SaveToDb(clientId string) error {
	dynamoItem := dynamodb.DynamoItem{
		ID:       clientId,
		DataType: "public-key",
		Data:     k.PublicKey,
		DataKind: "client-info",
	}

	return dynamodb.Put(dynamoItem)
}

func (k *KeyPair) Verify(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return nil, errors.New("Invalid claims")
		}
		issuer, ok := claims["iss"].(string)
		if !ok {
			return nil, errors.New("Invalid iss")
		}

		keyData, err := dynamodb.GetByIDDataType(issuer, "public-key")
		if err != nil {
			return nil, err
		}

		publicKeyBlock, rest := pem.Decode([]byte(keyData.Data))
		if publicKeyBlock == nil {
			return nil, errors.New("Public key block is null. rest: " + string(rest))
		}

		return x509.ParsePKCS1PublicKey(publicKeyBlock.Bytes)
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("Invalid token")
	}

	mapClaims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("Failed to parse claims")
	}

	if !mapClaims.VerifyExpiresAt(time.Now().Unix(), true) {
		return jwt.MapClaims{}, errors.New("Token expired")
	}

	return mapClaims, nil
}
