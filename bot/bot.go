package bot

import (
	"fmt"

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
