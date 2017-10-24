package main

import (
	"fmt"
	"strconv"

	"github.com/artberri/gitcleaner/services"
	"github.com/c2h5oh/datasize"
	"github.com/ryanuber/columnize"
	"github.com/urfave/cli"
)

func listCommand(c *cli.Context, git *services.GitManager) error {
	path, errNotRepoFound := git.EnsureRepoPath(c.Args().Get(0))
	if errNotRepoFound != nil {
		return errNotRepoFound
	}

	objects, errGettingObjects := getObjects(path, git)
	if errGettingObjects != nil {
		return errGettingObjects
	}

	output := []string{}
	max := c.Int("lines")
	hr := c.Bool("humanreadable")
	unique := c.Bool("unique")
	count := len(objects)

	if unique {
		objects = groupObjectsByFile(objects)
	}

	if max == 0 || count < max {
		max = count
	}
	for i := 0; i < max; i++ {
		var size string
		if hr {
			size = datasize.ByteSize(objects[i].size).HumanReadable()
		} else {
			size = strconv.FormatUint(objects[i].size, 10)
		}
		output = append(output, size+" | "+objects[i].path+" | "+objects[i].sha)
	}

	result := columnize.SimpleFormat(output)
	fmt.Println(result)

	return nil
}
