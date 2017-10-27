package domain

import (
	"bufio"
	"strings"
)

type MockFalseFileExister struct{}

func (r *MockFalseFileExister) Exists(path string) bool {
	return false
}

type MockTrueFileExister struct{}

func (r *MockTrueFileExister) Exists(path string) bool {
	return true
}

type MockRunner struct {
	Cmd string
}

func (r *MockRunner) Run(command string) (*bufio.Scanner, error) {
	r.Cmd = command
	return bufio.NewScanner(strings.NewReader("")), nil
}

type MockGitRepoManager struct {
	VerifyPackFn          func(path string) (*bufio.Scanner, error)
	VerifyPackInvoked     bool
	RevListFn             func(path string) (*bufio.Scanner, error)
	RevListInvoked        bool
	EnsureRepoPathFn      func(path string) (string, error)
	EnsureRepoPathInvoked bool
}

func (m *MockGitRepoManager) VerifyPack(path string) (*bufio.Scanner, error) {
	m.VerifyPackInvoked = true
	return m.VerifyPackFn(path)
}

// RevList executes git rev-list
func (m *MockGitRepoManager) RevList(path string) (*bufio.Scanner, error) {
	m.RevListInvoked = true
	return m.RevListFn(path)
}

// EnsureRepoPath ensure there is a Git repo
func (m *MockGitRepoManager) EnsureRepoPath(path string) (string, error) {
	m.EnsureRepoPathInvoked = true
	return m.EnsureRepoPathFn(path)
}
