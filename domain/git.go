package domain

import (
	"bufio"
	"errors"
)

// GitManager is a git repo manager
type GitManager struct {
	Runner  Runner
	Exister Exister
}

// VerifyPack executes git verify-pack
func (m *GitManager) VerifyPack(path string) (*bufio.Scanner, error) {
	gitCommand := "git verify-pack -v " + path + "/.git/objects/pack/pack-*.idx | egrep \"^\\w+ blob\\W+[0-9]+ [0-9]+ [0-9]+$\""

	return m.Runner.Run(gitCommand)
}

// RevList executes git rev-list
func (m *GitManager) RevList(path string) (*bufio.Scanner, error) {
	gitCommand := "git --git-dir=" + path + "/.git rev-list --all --objects"

	return m.Runner.Run(gitCommand)
}

// EnsureRepoPath ensure there is a Git repo
func (m *GitManager) EnsureRepoPath(path string) (string, error) {
	if path == "" {
		path = "."
	}

	if m.Exister.Exists(path + "/.git") {
		return path, nil
	}

	return path, errors.New("\"" + path + "\" is not a git repository path")
}
