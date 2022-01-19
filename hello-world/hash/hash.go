package hash

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"hello-world/config"

	"golang.org/x/crypto/sha3"
)

func CreateSha3_256Hash(val string) string {
	hash := sha3.New256()
	hash.Write([]byte(val))
	return string(hash.Sum(nil))
}

func CreateSha256HMAC(msg string) string {
	key := config.GetEnv("CHANNEL_SECRET")
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(msg))
	return hex.EncodeToString(mac.Sum(nil))
}
