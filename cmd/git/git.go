package git

import (
	"context"
	"fmt"

	"github.com/google/go-github/v37/github"
	"golang.org/x/oauth2"
)

var ctx = context.Background()

func CreateAuthenticatedClient(gitHubPat string) (*github.Client, error) {
	tokenSource := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: gitHubPat})
	tokenClient := oauth2.NewClient(ctx, tokenSource)

	client := github.NewClient(tokenClient)
	return client, nil
}

func CreatePullRequest(client *github.Client, owner string, repo string) error {
	title := "Test pull request"
	body := "Example description"
	head := "feature/test"
	base := "main"

	newPr, _, err := client.PullRequests.Create(ctx, owner, repo, &github.NewPullRequest{
		Title: &title,
		Body:  &body,
		Head:  &head,
		Base:  &base,
	})
	if err != nil {
		respErr := err.(*github.ErrorResponse)

		switch respErr.Response.StatusCode {
		case 422:
			fmt.Printf("\nPull request already exists. %s \n", respErr.Message)
			return err
		case 403:
			fmt.Printf("\nForbidden. %s \n", respErr.Message)
			return err
		default:
			fmt.Printf("\nError. %s \n", respErr.Message)
			return err
		}

	}

	fmt.Printf("Created pull request %d\n", newPr.GetNumber())

	return nil
}
