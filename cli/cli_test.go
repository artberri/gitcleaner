package cli

import (
	"os"
	"testing"
)

type MockListCommand struct {
	Path          string
	Max           int
	HumanReadable bool
	Unique        bool
}

// Exec executes the command
func (lc *MockListCommand) Exec(path string, max int, humanReadable bool, unique bool) error {
	lc.Path = path
	lc.Max = max
	lc.HumanReadable = humanReadable
	lc.Unique = unique

	return nil
}

func TestAppExecutesListCommandAndReadsPathAndDefaultValues(t *testing.T) {
	listCommand := &MockListCommand{}
	commands := Commands{
		List: listCommand,
	}
	app := App{}
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"cmd", "list", "/path/to/the/repo"}

	app.Start("0.0.1", commands)

	if listCommand.Path != "/path/to/the/repo" {
		t.Fatalf("Error executing list command: /path/to/the/repo path expected but got %s ", listCommand.Path)
	}

	if listCommand.HumanReadable != false {
		t.Fatalf("Error executing list command: false HumanReadable expected but got true ")
	}

	if listCommand.Max != 10 {
		t.Fatalf("Error executing list command: expected 10 as maximum default but got %d ", listCommand.Max)
	}

	if listCommand.Unique != false {
		t.Fatalf("Error executing list command: false Unique expected but got true ")
	}
}

func TestAppExecutesListCommandAndReadsMaximumArg(t *testing.T) {
	listCommand := &MockListCommand{}
	commands := Commands{
		List: listCommand,
	}
	app := App{}
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"cmd", "list", "/path/to/the/repo", "-n", "50"}

	app.Start("0.0.1", commands)

	if listCommand.Max != 50 {
		t.Fatalf("Error executing list command: 50 maximum expected but got %d ", listCommand.Max)
	}
}

func TestAppExecutesListCommandAndReadsMaximumLongArg(t *testing.T) {
	listCommand := &MockListCommand{}
	commands := Commands{
		List: listCommand,
	}
	app := App{}
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"cmd", "list", "/path/to/the/repo", "--lines", "50"}

	app.Start("0.0.1", commands)

	if listCommand.Max != 50 {
		t.Fatalf("Error executing list command: 50 maximum expected but got %d ", listCommand.Max)
	}
}

func TestAppExecutesListCommandAndReadsHumanReadableArg(t *testing.T) {
	listCommand := &MockListCommand{}
	commands := Commands{
		List: listCommand,
	}
	app := App{}
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"cmd", "list", "/path/to/the/repo", "-hr"}

	app.Start("0.0.1", commands)

	if listCommand.HumanReadable == false {
		t.Fatalf("Error executing list command: true HumanReadable expected but got false ")
	}
}

func TestAppExecutesListCommandAndReadsHumanReadableLongArg(t *testing.T) {
	listCommand := &MockListCommand{}
	commands := Commands{
		List: listCommand,
	}
	app := App{}
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"cmd", "list", "/path/to/the/repo", "--humanreadable"}

	app.Start("0.0.1", commands)

	if listCommand.HumanReadable == false {
		t.Fatalf("Error executing list command: true HumanReadable expected but got false ")
	}
}

func TestAppExecutesListCommandAndReadsUniqueArg(t *testing.T) {
	listCommand := &MockListCommand{}
	commands := Commands{
		List: listCommand,
	}
	app := App{}
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"cmd", "list", "/path/to/the/repo", "-u"}

	app.Start("0.0.1", commands)

	if listCommand.Unique == false {
		t.Fatalf("Error executing list command: true Unique expected but got false ")
	}
}

func TestAppExecutesListCommandAndReadsUniqueLongArg(t *testing.T) {
	listCommand := &MockListCommand{}
	commands := Commands{
		List: listCommand,
	}
	app := App{}
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"cmd", "list", "/path/to/the/repo", "--unique"}

	app.Start("0.0.1", commands)

	if listCommand.Unique == false {
		t.Fatalf("Error executing list command: true Unique expected but got false ")
	}
}
