package main

import (
	"bufio"
	"os/exec"
	"strings"

	"github.com/urfave/cli"
)

func verifyPack(path string) (*bufio.Scanner, *cli.ExitError) {
	gitCommand := "git verify-pack -v " + path + "/.git/objects/pack/pack-*.idx | egrep \"^\\w+ blob\\W+[0-9]+ [0-9]+ [0-9]+$\""

	cmd := exec.Command("bash", "-c", gitCommand)
	output, err := cmd.CombinedOutput()

	if err != nil {
		return nil, cli.NewExitError(err.Error(), 1)
	}

	return bufio.NewScanner(strings.NewReader(string(output))), nil
}

func revList(path string) (*bufio.Scanner, *cli.ExitError) {
	gitCommand := "git --git-dir=" + path + "/.git rev-list --all --objects"

	cmd := exec.Command("bash", "-c", gitCommand)
	output, err := cmd.CombinedOutput()

	if err != nil {
		return nil, cli.NewExitError(err.Error(), 1)
	}

	return bufio.NewScanner(strings.NewReader(string(output))), nil
}
