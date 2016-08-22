package github

import "fmt"

type Event struct {
	Action      string     `json:"action"`
	Issue       Issue      `json:"issue"`
	PullRequest Issue      `json:"pull_request"`
	Comment     Comment    `json:"comment"`
	Repository  Repository `json:"repository"`
}

func (e Event) String() string {
	return fmt.Sprintf("Title: %s, HEAD: %s, Base: %s,", e.Issue.Title, e.PullRequest.Head.Ref, e.PullRequest.Base.Ref)
}
