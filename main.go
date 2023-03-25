package main

import (
	"fmt"
	"os"

	"jack/git"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "jack",
		Usage: "A CLI for doing things",
		Commands: []*cli.Command{
			{
				Name:  "git",
				Usage: "Interact with git",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "pat",
						Aliases:  []string{"p"},
						Usage:    "Personal Access Token for Git authentication",
						Required: true,
					},
				},
				Action: func(c *cli.Context) error {
					fmt.Println("CREATING PULL REQUEST...")

					pat := c.String("pat")

					client, err := git.CreateAuthenticatedClient(pat)
					if err != nil {
						panic(err)
					}

					owner, repo := "rjackowens", "test-repo"

					err = git.CreatePullRequest(client, owner, repo)
					if err != nil {
						panic(err)
						// fmt.Println(err.Error())
					}
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}
