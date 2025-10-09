package types

type IssueCommentPayload struct {
	Action string `json:"action"`
	Issue  struct {
		Number      int `json:"number"`
		PullRequest struct {
			URL string `json:"url"`
		} `json:"pull_request"`
	} `json:"issue"`
	Comment struct {
		Body string `json:"body"`
		User struct {
			Login string `json:"login"`
		} `json:"user"`
	} `json:"comment"`
}
