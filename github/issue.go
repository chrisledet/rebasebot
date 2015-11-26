package github

type Issue struct {
	Body       string     `json:"body"`
	State      string     `json:"state"`
	Title      string     `json:"title"`
	Number     int        `json:"number"`
	Head       GitRef     `json:"head"`
	Base       GitRef     `json:"base"`
	Repository Repository `json:"repository"`
}
