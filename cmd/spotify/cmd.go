package spotify

import (
	"fmt"
	"strings"

	"github.com/urfave/cli/v2"
)

func SpotifyCommand() *cli.Command {

	// validMeetings := []string{"standup", "postmortem", "jourfix"}
	// meetings := cli.NewStringSlice(validMeetings...)

	return &cli.Command{
		Name:  "spotify",
		Usage: "Interact with spotify",
		Subcommands: []*cli.Command{
			{
				Name:  "login",
				Usage: "Login to Spotify",
				Flags: []cli.Flag{
					&cli.StringSliceFlag{
						Name:     "credentials",
						Aliases:  []string{"c"},
						Value:    cli.NewStringSlice("SPOTIFY_ID", "SPOTIFY_SECRET"),
						Usage:    "Spotify client credentials. Required if environment variable SPOTIFY_ID or SPOTIFY_SECRET not set",
						Required: false,
					},
				},
				Action: func(c *cli.Context) error {
					fmt.Println("Logging in to Spotify...")

					credentials := c.StringSlice("credentials")
					separatedArgs := c.Args() // by default --credentials expects multiple values wrapped in single string

					// if only one value was passed, check if it's a single string with space-separated values
					if len(credentials) == 1 && strings.Contains(credentials[0], " ") {
						credentials = strings.Split(credentials[0], " ")
					}

					// check for multiple separated arguments and add them to the credentials slice
					for _, arg := range separatedArgs.Slice() {
						credentials = append(credentials, arg)
					}

					LoginToSpotify(*cli.NewStringSlice(credentials...))

					return nil
				},
			},
		},
	}
}
