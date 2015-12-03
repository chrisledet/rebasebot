package integrations

import (
	"github.com/chrisledet/rebasebot/git"
)

// Ties the git operations together to perform a branch rebase
func GitRebase(repostioryPath, baseRef, headRef string) error {
	filepath := git.GetRepositoryFilePath(repostioryPath)
	remoteRepositoryURL := git.GenerateCloneURL(repostioryPath)

	if !git.Exists(filepath) {
		git.Clone(remoteRepositoryURL)
	}

	if err := git.Fetch(filepath); err != nil {
		return err
	}

	if err := git.Checkout(filepath, headRef); err != nil {
		return err
	}

	if err := git.Reset(filepath, headRef); err != nil {
		return err
	}

	if err := git.Rebase(filepath, baseRef); err != nil {
		return err
	}

	if err := git.Push(filepath, headRef); err != nil {
		return err
	}

	return nil
}
