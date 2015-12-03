package integrations

import (
	"github.com/chrisledet/rebasebot/git"
	"github.com/chrisledet/rebasebot/github"
)

// Ties the git operations together to perform a branch rebase
func GitRebase(pr github.PullRequest) error {
	filepath := git.GetRepositoryFilePath(pr.Repository.FullName)
	remoteRepositoryURL := git.GenerateCloneURL(pr.Repository.FullName)

	if !git.Exists(filepath) {
		git.Clone(remoteRepositoryURL)
	}

	if err := git.Fetch(filepath); err != nil {
		return err
	}

	if err := git.Checkout(filepath, pr.Head.Ref); err != nil {
		return err
	}

	if err := git.Reset(filepath, pr.Head.Ref); err != nil {
		return err
	}

	if err := git.Rebase(filepath, pr.Base.Ref); err != nil {
		return err
	}

	if err := git.Push(filepath, pr.Head.Ref); err != nil {
		return err
	}

	return nil
}
