package github

type GitRef struct {
	Ref string `json:"ref"`
	Sha string `json:"sha"`
}
