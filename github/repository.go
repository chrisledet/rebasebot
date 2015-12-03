package github

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Repository struct {
	FullName string `json:"full_name"`
	Name     string `json:"name"`
	GitUrl   string `json:"git_url"`
	SshUrl   string `json:"ssh_url"`
	Owner    User   `json:"owner"`
}

func (r Repository) FindPR(number int) (*PullRequest, error) {
	var pr PullRequest

	log.Println("github.find_pr.started")

	path := fmt.Sprintf("/repos/%s/pulls/%d", r.FullName, number)
	request := NewGitHubRequest(path)
	response, err := httpClient.Do(request)

	defer response.Body.Close()

	if err != nil {
		log.Println("github.find_pr.failed error: %s", err)
		return &pr, err
	}

	if response.StatusCode != http.StatusOK {
		log.Println("github.find_pr.failed status: ", response.StatusCode)
		return &pr, err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("github.find_pr.failed error:", err)
		return &pr, err
	}

	if err := json.Unmarshal(body, &pr); err != nil {
		log.Println("github.find_pr.failed error:", err)
		return &pr, err
	}

	log.Printf("github.find_pr.completed number: %d\n", pr.Number)

	return &pr, nil
}
