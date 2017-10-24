// gitcleaner - The Git Housekeeping Tool
// Copyright (C) 2017  Alberto Varela Sánchez <alberto@berriart.com>

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/gpl-3.0.html>.

package main

import (
	"os"

	"github.com/artberri/gitcleaner/services"
	"github.com/urfave/cli"
)

func main() {
	runner := &services.BashRunner{}
	exister := &services.FileExister{}
	git := &services.GitManager{
		Runner:  *runner,
		Exister: *exister,
	}

	app := cli.NewApp()
	app.Name = "gitcleaner"
	app.Version = "0.0.1"
	app.Usage = "Git Housekeeping Utility"

	app.Copyright = `
	gitcleaner - Copyright (C) 2017 Alberto Varela Sánchez

	This program comes with ABSOLUTELY NO WARRANTY.
	This is free software, and you are welcome to redistribute it
	under certain conditions. You should have received a copy of 
	the GNU General Public License along with this program.  
	If not, see <https://www.gnu.org/licenses/gpl-3.0.html>.
	`
	app.Authors = []cli.Author{cli.Author{Name: "Alberto Varela Sánchez", Email: "alberto@berriart.com"}}

	app.Commands = []cli.Command{
		{
			Name:      "list",
			Aliases:   []string{"l"},
			Usage:     "List heavier file objects in the repository history",
			ArgsUsage: "[/path/to/your/repo]",
			Action: func(c *cli.Context) error {
				return listCommand(c, git)
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
