package main

import (
	"discord-pr/bot"
	"discord-pr/config"
	"discord-pr/github"
	"discord-pr/mux"
	"fmt"
	"net/http"
)

func main() {
	bot, err := bot.New(config.BOT_TOKEN)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer bot.Close()
	fmt.Println("Bot is running...")

	mux := mux.New(bot)
	mux.Route("/", github.Webhook, github.GithubSignature)

	err = http.ListenAndServe(config.SERVER_ADDRESS, mux)
	if err != nil {
		fmt.Println("http server:", err)
		return
	}
}
