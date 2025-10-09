package events

import (
	"discord-pr/bot"
	"discord-pr/github/types"
	"encoding/json"
	"fmt"
)

func OnIssueComment(body []byte, bot *bot.Bot, forumID string) error {
	var payload types.IssueCommentPayload
	if err := json.Unmarshal(body, &payload); err != nil {
		return fmt.Errorf("invalid issue comment payload: %w", err)
	}

	if payload.Issue.PullRequest.URL == "" || payload.Action != "created" {
		return nil
	}

	title := fmt.Sprintf("PR #%d", payload.Issue.Number)
	message := fmt.Sprintf("💬 **%s** commented:\n> %s",
		payload.Comment.User.Login, payload.Comment.Body)

	id, err := bot.SearchThread(forumID, title, 0, true)
	if err != nil {
		return err
	}
	_, err = bot.ChannelMessageSend(id, message)
	return err
}
