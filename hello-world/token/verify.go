package auth

/* import (
	"hello-world/config"

	jwt "github.com/dgrijalva/jwt-go"
)

var channelSecret []byte = []byte(config.GetChannelSecret())

type JwtClaims struct {
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.StandardClaims
}

func VerifyToken(tokenString string) (JwtClaims, error) {
	jwt, err := jwt.ParseWithClaims(tokenString, &JwtClaims{}, func(t *jwt.Token) (interface{}, error) {
		return channelSecret, nil
	})
	if claims, ok := jwt.Claims.(*JwtClaims); ok && jwt.Valid {
		return *claims, nil
	} else {
		return *claims, err
	}
} */
