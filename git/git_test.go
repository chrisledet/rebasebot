package git

import "testing"

func TestExtractOrgFromURLWithGoodURL(t *testing.T) {
	expected := "chrisledet"
	result := extractOrgFromURL("git://github.com/chrisledet/dotfiles.git")

	if result != expected {
		t.Errorf("got %s, want: %s \n", result, expected)
	}
}

func TestExtractOrgFromURLWithBadURL(t *testing.T) {
	expected := "chrisledet"
	result := extractOrgFromURL("github.com/chrisledet/dotfiles.git")

	if result != expected {
		t.Errorf("got %s, want: %s \n", result, expected)
	}
}

func TestExtractRepoNameFromURLWithGoodURL(t *testing.T) {
	expected := "dotfiles"
	result := extractRepoNameFromURL("git://github.com/chrisledet/dotfiles.git")

	if result != expected {
		t.Errorf("got %s, want: %s \n", result, expected)
	}
}

func TestExtractRepoNameFromURLWithBadURL(t *testing.T) {
	expected := "dotfiles"
	result := extractRepoNameFromURL("github.com/chrisledet/dotfiles.git")

	if result != expected {
		t.Errorf("got %s, want: %s \n", result, expected)
	}
}

func TestExtractRepoNameFromURLWithoutExt(t *testing.T) {
	expected := "dotfiles"
	result := extractRepoNameFromURL("github.com/chrisledet/dotfiles")

	if result != expected {
		t.Errorf("got %s, want: %s \n", result, expected)
	}
}
