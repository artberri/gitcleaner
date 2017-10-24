package services

import (
	"bufio"
	"os/exec"
	"strings"

	"github.com/urfave/cli"
)

// Runner is a command runner
type Runner interface {
	Run(string) (*bufio.Scanner, *cli.ExitError)
}

// BashRunner is bash command runner
type BashRunner struct{}

// Run will execute a bash command
func (r BashRunner) Run(command string) (*bufio.Scanner, *cli.ExitError) {
	cmd := exec.Command("bash", "-c", command)
	output, err := cmd.CombinedOutput()

	if err != nil {
		return nil, cli.NewExitError(err.Error(), 1)
	}

	return bufio.NewScanner(strings.NewReader(string(output))), nil
}
