package config

import "os"

func GetChannelSecret() string {
	return os.Getenv("CHANNEL_SECRET")
}
