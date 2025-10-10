package events

import (
	"discord-pr/config"
	"discord-pr/github/types"
	"encoding/json"
	"fmt"

	"github.com/xySaad/gocord"
)

func OnPullRequest(body []byte, bot *gocord.Bot, channelID string) error {
	var payload types.PullRequestPayload
	if err := json.Unmarshal(body, &payload); err != nil {
		return fmt.Errorf("invalid PR payload: %s", err)
	}

	switch payload.Action {
	case "opened":
		postTitle := fmt.Sprintf("PR #%d: %s", payload.Number, payload.PullRequest.Title)
		postDescription := payload.Pretty(config.PR_NOTIFICATION_ROLE)

		tags, err := bot.GetTagIDs(channelID, "open", payload.PullRequest.Head.Ref)
		if err != nil {
			return fmt.Errorf("error getting tag ids: %w", err)
		}

		if err := bot.CreatePost(channelID, postTitle, postDescription, tags); err != nil {
			return fmt.Errorf("error creating forum post: %w", err)
		}

	case "closed":
		fmt.Printf(
			"Pull Request #%d on %s closed by %s\n",
			payload.Number,
			payload.Repository.Name,
			payload.PullRequest.User.Login,
		)
	}

	return nil
}
