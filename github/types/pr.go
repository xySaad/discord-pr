package types

import (
	"fmt"
	"strings"
)

type PullRequest struct {
	Title  string `json:"title"`
	Body   string `json:"body"`
	Merged bool   `json:"merged"`

	User struct {
		Login string `json:"login"`
	} `json:"user"`

	Head struct {
		Ref string `json:"ref"`
	} `json:"head"`
}

type Repository struct {
	Name  string `json:"name"`
	Owner struct {
		Login string `json:"login"`
	} `json:"owner"`
}

type PullRequestPayload struct {
	Action      string      `json:"action"`
	Number      int         `json:"number"`
	PullRequest PullRequest `json:"pull_request"`
	Repository  Repository  `json:"repository"`
}

func (payload *PullRequestPayload) Pretty() string {
	prURL := fmt.Sprintf(
		"https://github.com/%s/%s/pull/%d",
		payload.Repository.Owner.Login,
		payload.Repository.Name,
		payload.Number,
	)

	body := payload.PullRequest.Body
	if body == "" {
		body = "_No description provided._"
	} else {
		body = "> " + strings.ReplaceAll(body, "\n", "\n> ")
	}

	postDescription := fmt.Sprintf(
		"📦 **Repository:** %s\n"+
			"🌿 **Branch:** `%s`\n"+
			"👤 **Author:** %s\n"+
			"🔗 **Link:** %s\n"+
			"\n%s",
		payload.Repository.Name,
		payload.PullRequest.Head.Ref,
		payload.PullRequest.User.Login,
		prURL,
		body,
	)

	return postDescription
}
