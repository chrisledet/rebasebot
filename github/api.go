// Package github provides a simple client for the GitHub API
package github

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	ApiUrl string = "https://api.github.com"
)

var (
	username  string
	password  string
	signature string
)

func SetAuth(user string, pwd string) {
	username = user
	password = pwd
}

func SetSignature(sign string) {
	signature = sign
}

func Username() string {
	return username
}

func Signature() string {
	return signature
}

func FindPR(repo Repository, number int) PullRequest {
	var pr PullRequest

	log.Println("github.request.pr.started")

	fullUrl := fmt.Sprintf("%s/repos/%s/pulls/%d", ApiUrl, repo.FullName, number)

	client := &http.Client{}
	request, _ := http.NewRequest("GET", fullUrl, nil)
	request.SetBasicAuth(username, password)
	request.Header.Add("Accept", "application/vnd.github.v3+json")

	response, err := client.Do(request)
	defer response.Body.Close()

	if err != nil {
		log.Println("github.request.pr.failed:", err.Error())
		return pr
	}

	if response.StatusCode != http.StatusOK {
		log.Printf("github.request.pr.failed, status: %d\n", response.StatusCode)
		return pr
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("github.request.pr.failed:", err.Error())
		return pr
	}

	if err := json.Unmarshal(body, &pr); err != nil {
		log.Println("github.request.pr.failed:", err.Error())
		return pr
	}

	log.Printf("github.request.pr.completed: Found PR #%d\n", pr.Number)

	return pr
}
