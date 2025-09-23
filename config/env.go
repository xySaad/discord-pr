package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var _ = load_all()

var BOT_TOKEN = os.Getenv("BOT_TOKEN")
var SERVER_PORT = os.Getenv("SERVER_PORT")
var WEBHOOK_SECRET = os.Getenv("WEBHOOK_SECRET")
var FORUM_CHANNEL_ID = os.Getenv("FORUM_CHANNEL_ID")

func load_all() any {
	err := godotenv.Load()
	if err != nil {
		panic(fmt.Sprint("Error loading .env file:", err))
	}
	return nil
}
