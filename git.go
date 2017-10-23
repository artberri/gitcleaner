package main

import (
	"bufio"

	"github.com/urfave/cli"
)

func verifyPack(path string) (*bufio.Scanner, *cli.ExitError) {
	gitCommand := "git verify-pack -v " + path + "/.git/objects/pack/pack-*.idx | egrep \"^\\w+ blob\\W+[0-9]+ [0-9]+ [0-9]+$\""

	return runner.Run(gitCommand)
}

func revList(path string) (*bufio.Scanner, *cli.ExitError) {
	gitCommand := "git --git-dir=" + path + "/.git rev-list --all --objects"

	return runner.Run(gitCommand)
}
