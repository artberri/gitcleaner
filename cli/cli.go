package cli

import (
	"os"

	"github.com/urfave/cli"
)

// Commands is an struct that contains all app commands
type Commands struct {
	List Listcommand
}

// Listcommand is the interface for the list object command
type Listcommand interface {
	Exec(path string, max int, humanReadable bool, unique bool) error
}

// App is a cli command manager wrapper
type App struct{}

// Start will start the CLI app
func (a App) Start(version string, commands Commands) {
	app := cli.NewApp()
	app.Name = "gitcleaner"
	app.Version = version
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
				if err := commands.List.Exec(c.Args().Get(0), c.Int("lines"), c.Bool("humanreadable"), c.Bool("unique")); err != nil {
					return cli.NewExitError(err.Error(), 1)
				}

				return nil
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
