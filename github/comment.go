package github

type Comment struct {
	Body string `json:"body"`
	User User   `json:"user"`
}
