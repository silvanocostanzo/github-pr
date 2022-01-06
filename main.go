package main

import (
	"context"
	"fmt"
	"log"

	"github.com/google/go-github/v41/github"
)

type Checker interface {
	CheckRate() ([]*github.Repository, error)
}

type DefaultChecker struct {
	client github.Client
	ctx    context.Context
	user   string
}

func (dc *DefaultChecker) CheckRate() ([]*github.Repository, error) {
	repos, _, err := dc.client.Repositories.List(dc.ctx, dc.user, nil)
	return repos, err
}

func checkRateLimit(c Checker) ([]*github.Repository, error) {
	repos, err := c.CheckRate()
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
	// ctx := context.Background()
	// ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: "your token"})
	// tc := oauth2.NewClient(ctx, ts)
	// client := github.NewClient(tc)

	dc := &DefaultChecker{ctx: context.Background(), client: *github.NewClient(nil), user: "silvanocostanzo"}
	repos, err := checkRateLimit(dc)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(repos)

	// _, err := checkRateLimit(ctx, client)

	//
	// req, res, err := createNewPullRequest(ctx, client)
	//
	// if err != nil {
	// log.Fatal(err)
	// }
	//
	// fmt.Println(req)
	// fmt.Println(res)
}
