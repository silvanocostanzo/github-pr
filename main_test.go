package main

import (
	"log"
	"reflect"
	"testing"

	"github.com/google/go-github/v41/github"
)

func Test_checkRateLimit(t *testing.T) {
	t.Run("Rate limit no error", func(t *testing.T) {
		spyChecker := &SpyChecker{
			repo: []*github.Repository{{ID: github.Int64(3)}},
			err:  nil,
		}

		got, err := checkRateLimit(spyChecker)
		want := []*github.Repository{{ID: github.Int64(3)}}

		if err != nil {
			log.Fatal(err)
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("Rate limit error", func(t *testing.T) {
		spyChecker := &SpyChecker{
			repo: []*github.Repository{},
			err:  &github.RateLimitError{Message: "error rate limit"},
		}

		got, err := checkRateLimit(spyChecker)

		if got != nil {
			log.Fatal("an error is expected")
		}

		if err != spyChecker.err {
			log.Fatal("we have a problem")
		}
	})
}

type SpyChecker struct {
	repo []*github.Repository
	err  error
}

func (sc *SpyChecker) CheckRate() ([]*github.Repository, error) {
	return sc.repo, sc.err
}
