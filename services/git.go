package services

import (
	"bufio"

	"github.com/urfave/cli"
)

// RepoManager is an interface for git style repo managers
type RepoManager interface {
	VerifyPack(path string) (*bufio.Scanner, *cli.ExitError)
	RevList(path string) (*bufio.Scanner, *cli.ExitError)
	EnsureRepoPath(path string) (string, *cli.ExitError)
}

// GitManager is a git repo manager
type GitManager struct {
	Runner  Runner
	Exister Exister
}

// VerifyPack executes git verify-pack
func (m GitManager) VerifyPack(path string) (*bufio.Scanner, *cli.ExitError) {
	gitCommand := "git verify-pack -v " + path + "/.git/objects/pack/pack-*.idx | egrep \"^\\w+ blob\\W+[0-9]+ [0-9]+ [0-9]+$\""

	return m.Runner.Run(gitCommand)
}

// RevList executes git rev-list
func (m GitManager) RevList(path string) (*bufio.Scanner, *cli.ExitError) {
	gitCommand := "git --git-dir=" + path + "/.git rev-list --all --objects"

	return m.Runner.Run(gitCommand)
}

// EnsureRepoPath ensure there is a Git repo
func (m GitManager) EnsureRepoPath(path string) (string, *cli.ExitError) {
	if path == "" {
		path = "."
	}

	if m.Exister.Exists(path + "/.git") {
		return path, nil
	}

	return path, cli.NewExitError("\""+path+"\" is not a git repository path", 1)
}
