package main

import (
	"github.com/artberri/gitcleaner/cli"
	"github.com/artberri/gitcleaner/columnize"
	"github.com/artberri/gitcleaner/datasize"
	"github.com/artberri/gitcleaner/domain"
	"github.com/artberri/gitcleaner/exec"
	"github.com/artberri/gitcleaner/os"
)

func main() {
	const Version = "0.0.1"

	runner := &exec.BashRunner{}
	exister := &os.FileExister{}
	conv := &datasize.Converter{}
	col := &columnize.Columnizer{}
	git := &domain.GitManager{
		Runner:  runner,
		Exister: exister,
	}
	gom := &domain.GitObjectManager{
		Git: git,
	}
	listCommand := &domain.ListCommand{
		Converter:     conv,
		Columnizer:    col,
		ObjectManager: gom,
	}
	commands := cli.Commands{
		List: listCommand,
	}
	app := cli.App{}
	app.Start(Version, commands)
}
