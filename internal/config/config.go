package config

import (
	"os"
)

var (
	PORT           string
	REDIS_ADDRESS  string
	REDIS_PASSWORD string
)

func InitConfig() {
	PORT = os.Getenv("PORT")
	REDIS_ADDRESS = os.Getenv("REDIS_ADDRESS")
	REDIS_PASSWORD = os.Getenv("REDIS_PASSWORD")
}
