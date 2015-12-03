// Package github provides a simple client for the GitHub API
package github

import (
	"net/http"
	"strings"
)

const (
	mediaType   = "application/vnd.github.v3+json"
	contentType = "application/json"
	agent       = "rebasebot"
)

var (
	username   string
	password   string
	signature  string
	httpClient = &http.Client{}
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

// Returns a request set up for the GitHub API
func NewGitHubRequest(path string) *http.Request {
	requestUrl := "https://api.github.com" + path
	request, _ := http.NewRequest("GET", requestUrl, nil)
	request.SetBasicAuth(username, password)
	request.Header.Set("Accept", mediaType)
	request.Header.Set("Content-Type", contentType)
	request.Header.Set("User-Agent", agent)

	return request
}

// Check to see if logged in user was mentioned in comment
func WasMentioned(c Comment) bool {
	return strings.Contains(c.Body, "@"+username)
}
