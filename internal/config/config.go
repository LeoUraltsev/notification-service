package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Telegram struct {
		Token string
		Debug bool
	}
}

func New() *Config {
	godotenv.Load(".env")

	tkn := os.Getenv("TELEGRAM_BOT_TOKEN")
	deb := os.Getenv("TELEGRAM_BOT_DEBUG")
	tgDebug := false
	if deb == "true" {
		tgDebug = true
	}

	return &Config{
		Telegram: struct {
			Token string
			Debug bool
		}{
			Token: tkn,
			Debug: tgDebug,
		},
	}
}
