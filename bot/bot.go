package bot

import (
	"fmt"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

type Bot struct {
	*discordgo.Session
}

func New(token string) (*Bot, error) {
	session, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, fmt.Errorf("error creating bot: %s", err)
	}

	err = session.Open()
	if err != nil {
		return nil, fmt.Errorf("error opening connection: %s", err)
	}

	return &Bot{Session: session}, nil
}

// ThreadsList represents a list of threads alongisde with thread member objects for the current user.
type ThreadsList struct{ *discordgo.ThreadsList }

func (threads ThreadsList) SearchTitle(title string) (string, bool) {
	for _, t := range threads.Threads {
		if strings.HasPrefix(t.Name, title) {
			return t.ID, true
		}
	}

	return "", false
}

// ThreadsActive returns all active threads for specified channel.
func (b *Bot) ThreadsActive(channelID string, options ...discordgo.RequestOption) (*ThreadsList, error) {
	threads, err := b.Session.ThreadsActive(channelID)
	if err != nil {
		return nil, err
	}

	return &ThreadsList{threads}, nil
}

// ThreadsArchived returns archived threads for specified channel.
// before : If specified returns only threads before the timestamp
// limit  : Optional maximum amount of threads to return.
func (b *Bot) ThreadsArchived(channelID string, before *time.Time, limit int, options ...discordgo.RequestOption) (*ThreadsList, error) {
	threads, err := b.Session.ThreadsArchived(channelID, before, limit, options...)
	if err != nil {
		return nil, err
	}

	return &ThreadsList{threads}, nil
}
