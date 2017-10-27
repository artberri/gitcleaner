package domain

import (
	"testing"
)

func TestVerifyPackExecutesProperCommand(t *testing.T) {
	runner := &MockRunner{}
	exister := &MockTrueFileExister{}
	git := &GitManager{
		Runner:  runner,
		Exister: exister,
	}

	git.VerifyPack("./path/to/repo")

	if runner.Cmd != "git verify-pack -v ./path/to/repo/.git/objects/pack/pack-*.idx | egrep \"^\\w+ blob\\W+[0-9]+ [0-9]+ [0-9]+$\"" {
		t.Fatalf("Error on command: %s", runner.Cmd)
	}
}

func TestRevListExecutesProperCommand(t *testing.T) {
	runner := &MockRunner{}
	exister := &MockTrueFileExister{}
	git := &GitManager{
		Runner:  runner,
		Exister: exister,
	}

	git.RevList("./path/to/repo")

	if runner.Cmd != "git --git-dir=./path/to/repo/.git rev-list --all --objects" {
		t.Fatalf("Error on command: %s", runner.Cmd)
	}
}

func TestEnsureRepoPathReturnsInputWhenGitPathExists(t *testing.T) {
	runner := &MockRunner{}
	exister := &MockTrueFileExister{}
	git := &GitManager{
		Runner:  runner,
		Exister: exister,
	}

	repoPath, err := git.EnsureRepoPath("./path/to/repo")

	if repoPath != "./path/to/repo" {
		t.Fatalf("Not proper path returned: %s", repoPath)
	}

	if err != nil {
		t.Fatalf("Unexpected error thrown: %s", err.Error())
	}
}

func TestEnsureRepoPathReturnsDotWhenPathIsAnEmtyString(t *testing.T) {
	runner := &MockRunner{}
	exister := &MockTrueFileExister{}
	git := &GitManager{
		Runner:  runner,
		Exister: exister,
	}

	repoPath, err := git.EnsureRepoPath("")

	if repoPath != "." {
		t.Fatalf("Not proper path returned: %s", repoPath)
	}

	if err != nil {
		t.Fatalf("Unexpected error thrown: %s", err.Error())
	}
}

func TestEnsureRepoPathReturnsErrorIfItIsNotAGitRepo(t *testing.T) {
	runner := &MockRunner{}
	exister := &MockFalseFileExister{}
	git := &GitManager{
		Runner:  runner,
		Exister: exister,
	}

	_, err := git.EnsureRepoPath("./path/to/repo")

	if err == nil {
		t.Fatal("Error was expected")
	}
}
