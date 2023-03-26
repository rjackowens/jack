package spotify

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

func LoginToSpotify(credentials cli.StringSlice) (string, error) {

	if len(credentials.Value()) != 2 {
		fmt.Println("Both SPOTIFY_ID or SPOTIFY_SECRET must be passed")
		return "", fmt.Errorf("Oof")
	}

	clientID := os.Getenv(credentials.Value()[0])
	clientSecret := os.Getenv(credentials.Value()[1])

	if clientID == "" || clientSecret == "" {
		fmt.Println("SPOTIFY_ID or SPOTIFY_SECRET environment variables not set. Using --credentials")

		clientID := credentials.Value()[0]
		clientSecret := credentials.Value()[1]

		fmt.Printf("SPOTIFY_ID: %s\nSPOTIFY_SECRET: %s\n", clientID, clientSecret)
		return clientID, nil
	}

	fmt.Printf("SPOTIFY_ID_2: %s\nSPOTIFY_SECRET_2: %s\n", clientID, clientSecret)
	return clientID, nil

}
