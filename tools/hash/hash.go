package hash

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"os"

	"golang.org/x/crypto/sha3"
)

func CreateSha3_256Hash(val string) string {
	hash := sha3.New256()
	_, err := hash.Write([]byte(val))
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(hash.Sum(nil))
}

func CreateSha256HMAC(msg string) []byte {
	key := os.Getenv("CHANNEL_SECRET")
	mac := hmac.New(sha256.New, []byte(key))
	_, err := mac.Write([]byte(msg))
	if err != nil {
		panic(err)
	}
	return mac.Sum(nil)
}
