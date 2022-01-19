package cookie

import (
	"errors"
	"strings"
)

func GetCookieValue(cookie []string, key string) (string, error) {
	for _, v := range cookie {
		if strings.HasPrefix(v, key) {
			return strings.Split(v, "=")[1], nil
		}
	}
	return "", errors.New("Cookie not found")
}
