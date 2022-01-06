package main

import (
	"context"
	"fmt"
	"log"

	"github.com/google/go-github/v41/github"
)

type Checker interface {
	CheckRate() ([]*github.Repository, error)
	CheckRepo() (*github.Repository, error)
}

type DefaultChecker struct {
	client github.Client
	ctx    context.Context
	repo   string
	user   string
}

// dc.repo is not used in CheckRate
func (dc *DefaultChecker) CheckRate() ([]*github.Repository, error) {
	repos, _, err := dc.client.Repositories.List(dc.ctx, dc.user, nil)
	return repos, err
}

func (dc *DefaultChecker) CheckRepo() (*github.Repository, error) {
	repo, _, err := dc.client.Repositories.Get(dc.ctx, dc.user, dc.repo)
	return repo, err
}

func checkRateLimit(c Checker) ([]*github.Repository, error) {
	repos, err := c.CheckRate()
	if _, ok := err.(*github.RateLimitError); ok {
		log.Println("hit rate limit")
		return []*github.Repository{}, err
	}
	return repos, nil
}

func getRepo(c Checker) (*github.Repository, error) {
	repo, err := c.CheckRepo()
	if err != nil {
		log.Fatalf("cannot get the repo")
	}
	return repo, nil
}

func main() {

	dc := &DefaultChecker{ctx: context.Background(), client: *github.NewClient(nil), user: "silvanocostanzo", repo: "assmat"}
	repo, err := getRepo(dc)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(*repo.Name)

	_, err = checkRateLimit(dc)
	if err != nil {
		log.Fatal(err)
	}

}
