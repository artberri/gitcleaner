package exec

import (
	"bufio"
	"os/exec"
	"strings"

	"github.com/urfave/cli"
)

// BashRunner is bash command runner
type BashRunner struct{}

// Run will execute a bash command
func (r BashRunner) Run(command string) (*bufio.Scanner, error) {
	cmd := exec.Command("bash", "-c", command)
	output, err := cmd.CombinedOutput()

	if err != nil {
		return nil, cli.NewExitError(string(output), 1)
	}

	return bufio.NewScanner(strings.NewReader(string(output))), nil
}
