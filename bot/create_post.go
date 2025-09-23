package bot

import "github.com/bwmarrin/discordgo"

func (bot *Bot) CreateTag(channelID string, forumTag discordgo.ForumTag) (*discordgo.ForumTag, error) {
	ch, err := bot.Channel(channelID)
	if err != nil {
		return nil, err
	}

	// copy existing tags
	tags := append(ch.AvailableTags, forumTag)

	updated, err := bot.ChannelEditComplex(channelID, &discordgo.ChannelEdit{
		AvailableTags: &tags,
	})
	if err != nil {
		return nil, err
	}

	// return the last one (the new tag)
	return &updated.AvailableTags[len(updated.AvailableTags)-1], nil
}

// Fetches and returns tag IDs mapped by their names
func (bot *Bot) FetcheTags(channelID string) (map[string]string, error) {
	channel, err := bot.Channel(channelID)
	if err != nil {
		return nil, err
	}

	mappedTags := map[string]string{}
	for _, tag := range channel.AvailableTags {
		mappedTags[tag.Name] = tag.ID
	}

	return mappedTags, nil
}

// Fetches or creates tags and returns their ids
func (bot *Bot) GetTagIDs(channelID string, tags ...string) ([]string, error) {
	tagIDs, err := bot.FetcheTags(channelID)
	if err != nil {
		return nil, err
	}

	result := []string{}
	for _, tagName := range tags {
		id, exist := tagIDs[tagName]

		if exist {
			result = append(result, id)
			continue
		}

		newTag, err := bot.CreateTag(channelID, discordgo.ForumTag{Name: tagName})
		if err != nil {
			return nil, err
		}

		result = append(result, newTag.ID)
	}

	return result, nil
}

// Creates a forum post/thread
func (bot *Bot) CreatePost(forumID string, title string, content string, tags []string) error {
	_, err := bot.ForumThreadStartComplex(forumID, &discordgo.ThreadStart{
		Name:                title,
		AutoArchiveDuration: 1440,
		AppliedTags:         tags,
	}, &discordgo.MessageSend{
		Content: content,
	})

	return err
}
