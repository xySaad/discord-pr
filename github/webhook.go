package github

import (
	"discord-pr/bot"
	"discord-pr/config"
	"discord-pr/github/events"
	"discord-pr/mux"
	"fmt"
	"net/http"
)

func Webhook(w http.ResponseWriter, r *mux.Request, bot *bot.Bot) {
	h := r.Header.Get
	fmt.Println("GitHub Event:", h("X-GitHub-Event"))
	switch h("X-GitHub-Event") {
	case "pull_request":
		err := events.OnPullRequest(r.Body(), bot, config.FORUM_CHANNEL_ID)
		if err != nil {
			fmt.Println(err)
		}

	case "issue_comment":
		err := events.OnIssueComment(r.Body(), bot, config.FORUM_CHANNEL_ID)
		if err != nil {
			fmt.Println(err)
		}

	case "pull_request_review":
		err := events.OnPullRequestReview(r.Body(), bot, config.FORUM_CHANNEL_ID)
		if err != nil {
			fmt.Println(err)
		}
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
