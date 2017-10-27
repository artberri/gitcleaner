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

func (m *MockGitRepoManager) RevList(path string) (*bufio.Scanner, error) {
	m.RevListInvoked = true
	return m.RevListFn(path)
}

func (m *MockGitRepoManager) EnsureRepoPath(path string) (string, error) {
	m.EnsureRepoPathInvoked = true
	return m.EnsureRepoPathFn(path)
}

type MockObjectManager struct {
	GetFn                     func(path string) ([]GitObject, error)
	GetInvoked                bool
	GroupObjectsByFileFn      func(oldObjects []GitObject) []GitObject
	GroupObjectsByFileInvoked bool
}

func (m *MockObjectManager) Get(path string) ([]GitObject, error) {
	m.GetInvoked = true
	return m.GetFn(path)
}

func (m *MockObjectManager) GroupObjectsByFile(oldObjects []GitObject) []GitObject {
	m.GroupObjectsByFileInvoked = true
	return m.GroupObjectsByFileFn(oldObjects)
}

type MockConverter struct {
	HumanReadableFn      func(size uint64) string
	HumanReadableInvoked bool
}

func (m *MockConverter) HumanReadable(size uint64) string {
	m.HumanReadableInvoked = true
	return m.HumanReadableFn(size)
}

type MockColumnizer struct {
	ColumnizeFn      func(rows []string) string
	ColumnizeInvoked bool
}

func (m *MockColumnizer) Columnize(rows []string) string {
	m.ColumnizeInvoked = true
	return m.ColumnizeFn(rows)
}
