package spotify

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func SpotifyCommand() *cli.Command {
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

					credentials := cli.NewStringSlice(c.StringSlice("credentials")...)
					LoginToSpotify(*credentials)

					return nil
				},
			},
		},
	}
}
