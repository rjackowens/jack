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

	pr, _, err := client.PullRequests.Create(ctx, owner, repo, &github.NewPullRequest{
		Title: &title,
		Body:  &body,
		Head:  &head,
		Base:  &base,
	})
	if err != nil {
		return err
	}
	fmt.Printf("Created pull request %d\n", pr.GetNumber())
	return nil
}
