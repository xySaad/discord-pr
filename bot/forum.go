package bot

// func (bot *Bot) get_tag() {
// 	// Create forum post
// 	tags := []*discordgo.ForumTag{
// 		{Name: payload.Branch, Moderated: false},
// 		{Name: "open", Moderated: false},
// 	}

// 	// Check if tags already exist in the forum
// 	channel, err := bot.Channel(forumChannelID)
// 	if err != nil {
// 		return fmt.Errorf("error fetching forum channel: %s", err)
// 	}
// 	// Update or create tags
// 	availableTags := channel.AvailableTags
// 	for _, tag := range tags {
// 		found := false
// 		for _, existingTag := range availableTags {
// 			if existingTag.Name == tag.Name {
// 				tag.ID = existingTag.ID
// 				found = true
// 				break
// 			}
// 		}
// 		if !found {
// 			availableTags = append(availableTags, *tag)
// 		}
// 	}

// 	// Update forum channel with new tags if needed
// 	if len(availableTags) > len(channel.AvailableTags) {
// 		_, err = bot.ChannelEdit(forumChannelID, &discordgo.ChannelEdit{
// 			AvailableTags: &availableTags,
// 		})
// 		if err != nil {
// 			return fmt.Errorf("error updating forum tags: %s", err)
// 		}
// 	}

// 	appliedTagIDs := []string{}
// 	for _, tag := range availableTags {
// 		if tag.Name == payload.PullRequest.Head.Ref || tag.Name == "opened" {
// 			appliedTagIDs = append(appliedTagIDs, tag.ID)
// 		}
// 	}

// }
