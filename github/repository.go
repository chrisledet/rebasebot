package github

type Repository struct {
	FullName string `json:"full_name"`
	Name     string `json:"name"`
	GitUrl   string `json:"git_url"`
	SshUrl   string `json:"ssh_url"`
	Owner    User   `json:"owner"`
}
