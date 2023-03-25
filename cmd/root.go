package cmd

import (
	"fmt"
	"os"

	"jack/cmd/git"

	"github.com/urfave/cli/v2"
)

func RootCommand() {
	app := &cli.App{
		Name:  "jack",
		Usage: "A CLI for doing things",
		Commands: []*cli.Command{
			git.GitCommand(),
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}