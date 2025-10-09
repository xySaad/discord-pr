package types

type PullRequestReviewPayload struct {
	Action      string `json:"action"`
	PullRequest struct {
		Number int `json:"number"`
	} `json:"pull_request"`
	Review struct {
		State string `json:"state"`
		Body  string `json:"body"`
		User  struct {
			Login string `json:"login"`
		} `json:"user"`
	} `json:"review"`
}
