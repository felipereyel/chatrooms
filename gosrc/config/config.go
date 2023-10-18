package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

type tconfigs struct {
	PublicDir     string
	ServerAddress string
}

var Configs tconfigs

func init() {
	publicDir := os.Getenv("PUBLIC_DIR")
	if publicDir != "" {
		Configs.PublicDir = publicDir
	} else {
		Configs.PublicDir = "./dist"
	}

	envPort := os.Getenv("PORT")
	if envPort != "" {
		Configs.ServerAddress = ":" + envPort
	} else {
		Configs.ServerAddress = ":3000"
	}
}
