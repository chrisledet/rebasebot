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

// Clones the given Git url. Returns filepath to git repository
func Clone(repositoryUrl string) (string, error) {
	urlSplit := strings.Split(repositoryUrl, "/")
	repoNameWithExt := urlSplit[len(urlSplit)-1]
	orgName := urlSplit[len(urlSplit)-2]
	repoName := strings.Split(repoNameWithExt, ".")[0]
	repositoryPath := path.Join(repoParentDir, orgName, repoName)

	log.Println("git.clone.started:", repositoryPath)
	defer log.Println("git.clone.finished:", repositoryPath)

	cmd := exec.Command("git", "clone", repositoryUrl, repositoryPath)
	if err := cmd.Run(); err != nil {
		log.Println("git.clone.failed:", repositoryPath, err.Error())
		return "", err
	}

	return repositoryPath, nil
}

// Calls git fetch inside repository path
func Fetch(repositoryPath string) error {
	log.Println("git.fetch.started:", repositoryPath)
	defer log.Println("git.fetch.finished:", repositoryPath)

	cmd := exec.Command("git", "fetch", "origin")
	cmd.Dir = path.Join(".", repositoryPath)
	if err := cmd.Run(); err != nil {
		log.Println("git.fetch.failed:", repositoryPath, err.Error())
		return err
	}

	return nil
}

// Checks out a given git ref inside repository path
func Checkout(repositoryPath, gitRef string) error {
	log.Println("git.checkout.started:", repositoryPath, gitRef)
	defer log.Println("git.checkout.finished:", repositoryPath, gitRef)

	cmd := exec.Command("git", "checkout", gitRef)
	cmd.Dir = path.Join(".", repositoryPath)
	if err := cmd.Run(); err != nil {
		log.Println("git.checkout.failed:", repositoryPath, err.Error())
		return err
	}

	return nil
}

// Does hard reset inside repository path
func Reset(repositoryPath, branch string) error {
	log.Println("git.reset.started:", repositoryPath, branch)
	defer log.Println("git.reset.finished:", repositoryPath, branch)

	cmd := exec.Command("git", "reset", "--hard", branch)
	cmd.Dir = path.Join(".", repositoryPath)
	if err := cmd.Run(); err != nil {
		log.Println("git.reset.failed:", repositoryPath, err.Error())
		return err
	}

	return nil
}

// Rebases branch with baseBranch inside repository path
func Rebase(repositoryPath, baseBranch string) error {
	log.Println("git.rebase.started:", repositoryPath, baseBranch)
	defer log.Println("git.rebase.finished:", repositoryPath, baseBranch)

	cmd := exec.Command("git", "rebase", baseBranch)
	cmd.Dir = path.Join(".", repositoryPath)

	if err := cmd.Run(); err != nil {
		log.Printf("git.rebase.abort.started repo: %s, err: %s \n", repositoryPath, err.Error())

		abortCmd := exec.Command("git", "rebase", "--abort")
		abortCmd.Dir = path.Join(".", repositoryPath)

		if err := abortCmd.Run(); err != nil {
			log.Println("git.rebase.abort.failed:", repositoryPath)
		} else {
			log.Println("git.rebase.abort.finished:", repositoryPath)
		}

		return err
	}

	return nil
}

// Pushes branch back to origin
func Push(repositoryPath, branch string) error {
	log.Println("git.push.started:", repositoryPath, branch)
	defer log.Println("git.push.finished:", repositoryPath, branch)

	cmd := exec.Command("git", "push", "--force", "origin", branch)
	cmd.Dir = path.Join(".", repositoryPath)
	if err := cmd.Run(); err != nil {
		log.Println("git.push.failed:", repositoryPath, err.Error())
		return err
	}

	return nil
}

func Clean() {
	log.Printf("git.cache.cleanup.started: path: %s\n", repoParentDir)

	if err := os.RemoveAll(repoParentDir); err != nil {
		log.Fatalf("git.cache.cleanup.failed: path: %s, error: %s\n", repoParentDir, err.Error())
	}

	log.Printf("git.cache.cleanup.finished")
}
