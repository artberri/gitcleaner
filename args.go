package main

import (
	"os"

	"github.com/urfave/cli"
)

func getRepoPath(c *cli.Context) (string, *cli.ExitError) {
	var exitError *cli.ExitError
	repoPath := c.Args().Get(0)

	if repoPath == "" {
		repoPath = "."
	}

	_, err := os.Stat(repoPath + "/.git")

	if os.IsNotExist(err) {
		exitError = cli.NewExitError("\""+repoPath+"\" is not a git repository path", 1)
	}

	return repoPath, exitError
}
