package main

import (
	"discord-pr/config"
	"discord-pr/github"
	"fmt"
	"net/http"

	"github.com/xySaad/gocord"

	"github.com/xySaad/trail"
)

func main() {
	bot, err := gocord.New(config.BOT_TOKEN)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer bot.Close()
	fmt.Println("Bot is running...")

	router := trail.New(bot)
	router.Add("GET /", github.Webhook, github.GithubSignature)

	err = http.ListenAndServe(config.SERVER_ADDRESS, router)
	if err != nil {
		fmt.Println("http server:", err)
		return
	}
}
