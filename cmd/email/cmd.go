package email

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func EmailCommand() *cli.Command {
	return &cli.Command{
		Name:  "email",
		Usage: "Send emails",
		Subcommands: []*cli.Command{
			{
				Name:  "gmail",
				Usage: "Send an email via gmail",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "subject",
						Usage:    "Email Subject",
						Required: false,
					},
					&cli.StringFlag{
						Name:     "recipient",
						Usage:    "Person to email",
						Value:    "email@gmail.com",
						Required: false,
					},
					&cli.StringFlag{
						Name:     "body",
						Usage:    "Email body",
						Value:    "Test email",
						Required: false,
					},
				},
				Action: func(c *cli.Context) error {
					fmt.Println("Sending gmail email...")
					SendGmailEmail()
					return nil
				},
			},
		},
	}
}
