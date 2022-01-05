package main

import (
	"context"
	"fmt"
	"log"

	"github.com/google/go-github/v41/github"
	"golang.org/x/oauth2"
)

func checkRateLimit(ctx context.Context, client *github.Client) ([]*github.Repository, error) {
	repos, _, err := client.Repositories.List(ctx, "silvanocostanzo", nil)
	if _, ok := err.(*github.RateLimitError); ok {
		log.Println("hit rate limit")
		return nil, err
	}
	return repos, nil

}

func createNewPullRequest(ctx context.Context, client *github.Client) (*github.PullRequest, *github.Response, error) {
	newPR := &github.NewPullRequest{
		Title: github.String("test PR"),
		Head:  github.String("dev"),
		Base:  github.String("master"),
		Body:  github.String("a description"),
		// Issue:               github.Int(35),
		MaintainerCanModify: github.Bool(true),
		Draft:               github.Bool(false),
	}

	req, res, err := client.PullRequests.Create(ctx, "silvanocostanzo", "assmat", newPR)

	if err != nil {
		return nil, nil, err
	}

	return req, res, nil
}

func main() {
	ctx := context.Background()

	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: "your token"})
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	_, err := checkRateLimit(ctx, client)

	if err != nil {
		log.Fatal(err)
	}

	req, res, err := createNewPullRequest(ctx, client)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(req)
	fmt.Println(res)
}
