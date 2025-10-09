package events

import (
	"discord-pr/bot"
	"discord-pr/github/types"
	"encoding/json"
	"fmt"
)

func OnPullRequestReview(body []byte, bot *bot.Bot, forumID string) error {
	var payload types.PullRequestReviewPayload
	if err := json.Unmarshal(body, &payload); err != nil {
		return fmt.Errorf("invalid PR review payload: %w", err)
	}

	if payload.Review.State != "changes_requested" {
		return nil
	}

	threadPrefix := fmt.Sprintf("PR #%d", payload.PullRequest.Number)
	message := fmt.Sprintf("🧾 **%s** requested changes:\n> %s",
		payload.Review.User.Login, payload.Review.Body)

	id, err := bot.SearchThread(forumID, threadPrefix, 0)
	if err != nil {
		return err
	}
	_, err = bot.ChannelMessageSend(id, message)
	return err
}
