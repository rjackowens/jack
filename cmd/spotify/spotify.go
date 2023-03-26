package spotify

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

func LoginToSpotify(credentials cli.StringSlice) (string, error) {

	// if len(credentials.Value()) != 2 {
	// 	return "", fmt.Errorf("invalid number of credentials")
	// }
	clientID := os.Getenv(credentials.Value()[0])
	clientSecret := os.Getenv(credentials.Value()[1])

	if clientID == "" || clientSecret == "" {
		// log.Fatal("SPOTIFY_ID or SPOTIFY_SECRET not set")
		fmt.Println("SPOTIFY_ID or SPOTIFY_SECRET environment variables not set. Using --credentials")
		fmt.Println("CREDENTIALS:", credentials)

		if credentials.Value()[0] == "" || credentials.Value()[1] == "" {
			// return "", fmt.Errorf("invalid number of credentials")
			fmt.Println("Invalid number of values passed to credentials")
		}

	}

	// fmt.Printf("SPOTIFY_ID: %s\nSPOTIFY_SECRET: %s\n", clientID, clientSecret)
	return "", nil

}
