// Package github provides a simple client for the GitHub API
package github

import (
	"net/http"
)

const (
	mediaType = "application/vnd.github.v3+json"
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
func NewGitHubRequest(path string, httpMethod string) *http.Request {
	requestUrl := "https://api.github.com" + path
	request, _ := http.NewRequest(httpMethod, requestUrl, nil)
	request.SetBasicAuth(username, password)
	request.Header.Add("Accept", mediaType)

	return request
}
