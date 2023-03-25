package git

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func GitCommand() *cli.Command {
	return &cli.Command{
		Name:  "git",
		Usage: "Interact with git",
		Subcommands: []*cli.Command{
			{
				Name:  "pr",
				Usage: "Create a pull request",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "pat",
						Usage:    "Personal Access Token for authentication",
						Required: true,
					},
					&cli.StringFlag{
						Name:     "owner",
						Usage:    "Owner of the repository",
						Value:    "rjackowens",
						Required: false,
					},
					&cli.StringFlag{
						Name:     "repo",
						Usage:    "Repository name",
						Value:    "test-repo",
						Required: false,
					},
					&cli.StringFlag{
						Name:     "title",
						Usage:    "Title of the PR",
						Value:    "Default Title",
						Required: false,
					},
					&cli.StringFlag{
						Name:     "branch",
						Usage:    "Target branch to merge PR into",
						Value:    "main",
						Required: false,
					},
				},
				Action: func(c *cli.Context) error {
					fmt.Println("CREATING PULL REQUEST...")

					client, err := CreateAuthenticatedClient(c.String("pat"))
					if err != nil {
						panic(err)
					}

					err = CreatePullRequest(client, c.String("owner"), c.String("repo"))
					if err != nil {
						panic(err)
						// fmt.Println(err.Error())
					}
					return nil
				},
			},
		},
	}
}
