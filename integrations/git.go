package integrations

import (
	"path"

	"github.com/chrisledet/rebasebot/git"
	"github.com/chrisledet/rebasebot/github"
)

// Ties the git operations together to perform a branch rebase
func GitRebase(pr *github.PullRequest) error {
	
	filepath := git.GetRepositoryFilePath(pr.Head.Repository.FullName)
	remoteRepositoryURL := git.GenerateCloneURL(pr.Head.Repository.FullName)

	if !git.Exists(filepath) {
		if _, err := git.Clone(remoteRepositoryURL); err != nil {
			pr.PostComment("I could not pull " + pr.Head.Repository.FullName + " from GitHub.")
			return err
		}
	}

	if err := git.Fetch(filepath); err != nil {
		git.Prune(filepath)
		pr.PostComment("I could not fetch the latest changes from GitHub. Please try again in a few minutes.")
		return err
	}

	if err := git.Checkout(filepath, pr.Head.Ref); err != nil {
		pr.PostComment("I could not checkout " + pr.Head.Ref + " locally.")
		return err
	}

	if err := git.Reset(filepath, path.Join("origin", pr.Head.Ref)); err != nil {
		pr.PostComment("I could not checkout " + pr.Head.Ref + " locally.")
		return err
	}

	if err := git.Config(filepath, "user.name", git.GetName()); err != nil {
		pr.PostComment("I could run git config for user.name on the server.")
		return err
	}

	if err := git.Config(filepath, "user.email", git.GetEmail()); err != nil {
		pr.PostComment("I could run git config for user.email on the server.")
		return err
	}

	if err := git.Rebase(filepath, path.Join("origin", pr.Base.Ref)); err != nil {
		pr.PostComment("I could not rebase " + pr.Head.Ref + " with " + pr.Base.Ref + ". There are conflicts.")
		return err
	}

	if err := git.Push(filepath, pr.Head.Ref); err != nil {
		pr.PostComment("I could not push the changes to " + pr.Base.Ref + ".")
		return err
	}

	pr.PostComment("I just pushed up the changes, enjoy!")
	return nil
}
