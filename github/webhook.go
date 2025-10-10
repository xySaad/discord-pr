package github

import (
	"discord-pr/config"
	"discord-pr/github/events"
	"fmt"
	"net/http"

	"github.com/xySaad/gocord"

	"github.com/xySaad/trail"
)

func Webhook(c *trail.Context[*gocord.Bot]) {
	fmt.Println("GitHub Event:", c.Header("X-GitHub-Event"))
	switch c.Header("X-GitHub-Event") {
	case "pull_request":
		err := events.OnPullRequest(c.BodyNoErr(), c.Dep, config.FORUM_CHANNEL_ID)
		if err != nil {
			fmt.Println(err)
		}

	case "issue_comment":
		err := events.OnIssueComment(c.BodyNoErr(), c.Dep, config.FORUM_CHANNEL_ID)
		if err != nil {
			fmt.Println(err)
		}

	case "pull_request_review":
		err := events.OnPullRequestReview(c.BodyNoErr(), c.Dep, config.FORUM_CHANNEL_ID)
		if err != nil {
			fmt.Println(err)
		}
	}

	c.Response.WriteHeader(http.StatusOK)
	c.Write([]byte("OK"))
}
