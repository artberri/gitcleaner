package services

import (
	"os"
)

// Exister is an interface for existence checkers
type Exister interface {
	Exists(string) bool
}

// FileExister is a file stater
type FileExister struct{}

// Exists will check if a file exists
func (r FileExister) Exists(path string) bool {
	_, err := os.Stat(path)

	return !os.IsNotExist(err)
}
