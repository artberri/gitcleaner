package main

import (
	"bufio"
	"strings"
	"testing"

	"github.com/urfave/cli"
)

type TestRunner struct{}

var cmd string

func (r TestRunner) Run(command string) (*bufio.Scanner, *cli.ExitError) {
	cmd = command
	return bufio.NewScanner(strings.NewReader("")), nil
}

func TestVerifyPackExecutesProperCommand(t *testing.T) {
	runner = TestRunner{}

	verifyPack("./path/to/repo")

	if cmd != "git verify-pack -v ./path/to/repo/.git/objects/pack/pack-*.idx | egrep \"^\\w+ blob\\W+[0-9]+ [0-9]+ [0-9]+$\"" {
		t.Fatalf("Error on command: %s", cmd)
	}
}

func TestRevListExecutesProperCommand(t *testing.T) {
	runner = TestRunner{}

	revList("./path/to/repo")

	if cmd != "git --git-dir=./path/to/repo/.git rev-list --all --objects" {
		t.Fatalf("Error on command: %s", cmd)
	}
}
