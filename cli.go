package main

import (
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "gitcleaner"
	app.Usage = "git housekeeping utility"

	app.Commands = []cli.Command{
		{
			Name:      "list",
			Aliases:   []string{"l"},
			Usage:     "List heavier file objects in the repository history",
			ArgsUsage: "[/path/to/your/repo]",
			Action: func(c *cli.Context) error {
				return listCommand(c)
			},
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "humanreadable, hr",
					Usage: "Outputs the size in a readable format",
				},
				cli.BoolFlag{
					Name:  "unique, u",
					Usage: "Outputs the size of the whole history grouped by file path",
				},
				cli.IntFlag{
					Name:  "lines, n",
					Usage: "Output a maximum of `NUM` files/objects, 0 = no limit",
					Value: 10,
				},
			},
		},
	}

	app.Run(os.Args)
}
