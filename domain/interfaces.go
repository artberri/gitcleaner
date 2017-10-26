package domain

import (
	"bufio"
)

// ObjectManager is an interface for object management
type ObjectManager interface {
	Get(path string) ([]GitObject, error)
	GroupObjectsByFile(oldObjects []GitObject) []GitObject
}

// ErrorFactory error factory
type ErrorFactory interface {
	New(err string) error
}

// RepoManager is an interface for git style repo managers
type RepoManager interface {
	VerifyPack(path string) (*bufio.Scanner, error)
	RevList(path string) (*bufio.Scanner, error)
	EnsureRepoPath(path string) (string, error)
}

// Exister is an interface for existence checkers
type Exister interface {
	Exists(path string) bool
}

// Runner is a command runner
type Runner interface {
	Run(string) (*bufio.Scanner, error)
}

// Converter converts byte sizes
type Converter interface {
	HumanReadable(size uint64) string
}

// Columnizer create columnes in plain text
type Columnizer interface {
	Columnize(rows []string) string
}
