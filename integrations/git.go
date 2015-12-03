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
		pr.PostComment("I could not fetch the latest changes from GitHub.")
		return err
	}

	if err := git.Checkout(filepath, pr.Head.Ref); err != nil {
		pr.PostComment("I could not checkout " + pr.Head.Ref + " locally.")
		return err
	}

	if err := git.Reset(filepath, pr.Head.Ref); err != nil {
		pr.PostComment("I could not checkout " + pr.Head.Ref + " locally.")
		return err
	}

	if err := git.Rebase(filepath, pr.Base.Ref); err != nil {
		pr.PostComment("I could not rebase your PR with " + pr.Base.Ref + ". There were conflicts.")
		return err
	}

	if err := git.Push(filepath, pr.Head.Ref); err != nil {
		pr.PostComment("I could not push to " + pr.Base.Ref + ".")
		return err
	}

	pr.PostComment("I just pushed up the changes, enjoy!")
	return nil
}
