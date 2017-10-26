package domain

import (
	"bufio"
	"strings"
	"testing"

	"github.com/urfave/cli"
)

var cmd string

type TestRunner struct{}

func (r TestRunner) Run(command string) (*bufio.Scanner, *cli.ExitError) {
	cmd = command
	return bufio.NewScanner(strings.NewReader("")), nil
}

type FalseFileExister struct{}

func (r FalseFileExister) Exists(path string) bool {
	return false
}

type TrueFileExister struct{}

func (r TrueFileExister) Exists(path string) bool {
	return true
}

func TestVerifyPackExecutesProperCommand(t *testing.T) {
	runner := &TestRunner{}
	exister := &TrueFileExister{}
	git := &GitManager{
		Runner:  runner,
		Exister: exister,
	}

	git.VerifyPack("./path/to/repo")

	if cmd != "git verify-pack -v ./path/to/repo/.git/objects/pack/pack-*.idx | egrep \"^\\w+ blob\\W+[0-9]+ [0-9]+ [0-9]+$\"" {
		t.Fatalf("Error on command: %s", cmd)
	}
}

func TestRevListExecutesProperCommand(t *testing.T) {
	runner := &TestRunner{}
	exister := &TrueFileExister{}
	git := &GitManager{
		Runner:  runner,
		Exister: exister,
	}

	git.RevList("./path/to/repo")

	if cmd != "git --git-dir=./path/to/repo/.git rev-list --all --objects" {
		t.Fatalf("Error on command: %s", cmd)
	}
}

func TestEnsureRepoPathReturnsInputWhenGitPathExists(t *testing.T) {
	runner := &TestRunner{}
	exister := &TrueFileExister{}
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
	runner := &TestRunner{}
	exister := &TrueFileExister{}
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
	runner := &TestRunner{}
	exister := &FalseFileExister{}
	git := &GitManager{
		Runner:  runner,
		Exister: exister,
	}

	_, err := git.EnsureRepoPath("./path/to/repo")

	if err == nil {
		t.Fatal("Error was expected")
	}
}
