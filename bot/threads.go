package bot

import (
	"fmt"
	"time"
)

func (b *Bot) SearchThread(forumID string, title string, archivedLimit int) (string, error) {
	threadsList, err := b.ThreadsActive(forumID)
	if err != nil {
		return "", fmt.Errorf("failed to list threads: %w", err)
	}

	if id, found := threadsList.SearchTitle(title); found {
		return id, nil
	}

	before := time.Now()
	hasMore := true

	for hasMore {
		threadsList, err = b.ThreadsArchived(forumID, &before, archivedLimit)
		if err != nil {
			return "", fmt.Errorf("failed to list threads: %w", err)
		}
		if id, found := threadsList.SearchTitle(title); found {
			return id, nil
		}

		hasMore = threadsList.HasMore
		threads := threadsList.Threads
		before = threads[len(threads)-1].ThreadMetadata.ArchiveTimestamp
	}

	return "", fmt.Errorf("thread named %s not found", title)
}
