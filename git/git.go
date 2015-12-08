// Package git provides basic git client functionality
package git

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"
)

var (
	repoParentDir string
	username      string
	password      string
)

type Output struct {
	Buffer string
}

func (w *Output) Write(b []byte) (int, error) {
	w.Buffer = w.Buffer + string(b)

	return len(b), nil
}

func (o *Output) String() string {
	return o.Buffer
}

func init() {
	repoParentDir = os.TempDir()
}

func SetParentDir(path string) {
	repoParentDir = path
}

func ParentDir() string {
	return repoParentDir
}

func SetAuth(user string, pwd string) {
	username = user
	password = pwd
}

func GenerateCloneURL(repositoryFullName string) string {
	return fmt.Sprintf("https://%s:%s@github.com/%s.git", username, password, repositoryFullName)
}

func Exists(repositoryPath string) bool {
	_, err := os.Stat(repositoryPath)
	return !os.IsNotExist(err)
}

func GetRepositoryFilePath(name string) string {
	return path.Join(repoParentDir, name)
}

// Clone executes a git clone command on the system. It returns the path to the repository on the system.
func Clone(repositoryUrl string) (string, error) {
	orgName := extractOrgFromURL(repositoryUrl)
	repoName := extractRepoNameFromURL(repositoryUrl)
	repositoryPath := path.Join(repoParentDir, orgName, repoName)

	log.Println("git.clone.started:", repositoryPath)

	cmd := exec.Command("git", "clone", repositoryUrl, repositoryPath)
	if err := cmd.Run(); err != nil {
		log.Println("git.clone.failed:", repositoryPath, err.Error())
		return "", err
	}

	log.Println("git.clone.finished:", repositoryPath)

	return repositoryPath, nil
}

// Calls git fetch inside repository path
func Fetch(repositoryPath string) error {
	log.Println("git.fetch.started:", repositoryPath)

	cmd := exec.Command("git", "fetch", "origin")
	cmd.Dir = path.Join(repositoryPath)
	if err := cmd.Run(); err != nil {
		log.Println("git.fetch.failed:", repositoryPath, err.Error())
		return err
	}

	log.Println("git.fetch.finished:", repositoryPath)

	return nil
}

// Checks out a given git ref inside repository path
func Checkout(repositoryPath, gitRef string) error {
	log.Println("git.checkout.started:", repositoryPath, gitRef)

	cmd := exec.Command("git", "checkout", gitRef)
	cmd.Dir = path.Join(repositoryPath)
	if err := cmd.Run(); err != nil {
		log.Println("git.checkout.failed:", repositoryPath, err.Error())
		return err
	}

	log.Println("git.checkout.finished:", repositoryPath, gitRef)

	return nil
}

// Does hard reset inside repository path
func Reset(repositoryPath, branch string) error {
	log.Println("git.reset.started:", repositoryPath, branch)

	cmd := exec.Command("git", "reset", "--hard", branch)
	cmd.Dir = path.Join(repositoryPath)
	if err := cmd.Run(); err != nil {
		log.Println("git.reset.failed:", repositoryPath, err.Error())
		return err
	}

	log.Println("git.reset.finished:", repositoryPath, branch)

	return nil
}

// Rebases branch with baseBranch inside repository path
func Rebase(repositoryPath, baseBranch string) error {
	cmdOutput := &Output{Buffer: ""}
	log.Println("git.rebase.started:", repositoryPath, baseBranch)

	cmd := exec.Command("git", "rebase", baseBranch)
	cmd.Dir = path.Join(repositoryPath)
	cmd.Stderr = cmdOutput
	cmd.Stdout = cmdOutput

	if err := cmd.Run(); err != nil {
		log.Printf("git.rebase.failed repo: %s, err: %s \n", repositoryPath, err.Error())

		log.Printf("git.rebase.abort.started repo: %s, err: %s, stderr: %s \n", repositoryPath, err.Error(), cmdOutput.String())

		abortCmd := exec.Command("git", "rebase", "--abort")
		abortCmd.Dir = path.Join(repositoryPath)

		if err := abortCmd.Run(); err != nil {
			log.Println("git.rebase.abort.failed:", repositoryPath)
		} else {
			log.Println("git.rebase.abort.finished:", repositoryPath)
		}

		return err
	}

	log.Println("git.rebase.finished:", repositoryPath, baseBranch)

	return nil
}

// Pushes branch back to origin
func Push(repositoryPath, branch string) error {
	log.Println("git.push.started:", repositoryPath, branch)

	cmd := exec.Command("git", "push", "--force", "origin", branch)
	cmd.Dir = path.Join(repositoryPath)
	if err := cmd.Run(); err != nil {
		log.Println("git.push.failed:", repositoryPath, err.Error())
		return err
	}

	log.Println("git.push.finished:", repositoryPath, branch)

	return nil
}

func Clean() {
	log.Printf("git.cache.cleanup.started: path: %s\n", repoParentDir)

	if err := os.RemoveAll(repoParentDir); err != nil {
		log.Fatalf("git.cache.cleanup.failed: path: %s, error: %s\n", repoParentDir, err.Error())
	}

	log.Printf("git.cache.cleanup.finished")
}

func extractOrgFromURL(githubURL string) string {
	splitBySlash := strings.Split(githubURL, "/")
	return splitBySlash[len(splitBySlash)-2]
}

func extractRepoNameFromURL(githubURL string) string {
	splitBySlash := strings.Split(githubURL, "/")
	repoNameWithExt := splitBySlash[len(splitBySlash)-1]
	return strings.Split(repoNameWithExt, ".")[0]
}
