package main

import (
	"reflect"
	"testing"

	"github.com/google/go-github/v41/github"
)

func Test_checkRateLimit(t *testing.T) {
	tests := []struct {
		name       string
		spyChecker *spyChecker
		want       []*github.Repository
	}{
		{
			name: "Rate limit no error",
			spyChecker: &spyChecker{
				repo: []*github.Repository{{ID: github.Int64(3)}},
				err:  nil,
			},
			want: []*github.Repository{{ID: github.Int64(3)}},
		},
		{
			name: "Rate limit error",
			spyChecker: &spyChecker{
				repo: []*github.Repository{},
				err:  &github.RateLimitError{Message: "error rate limit"},
			},
			want: []*github.Repository{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			got, _ := checkRateLimit(test.spyChecker)
			want := test.want

			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	}
}

type spyChecker struct {
	repo []*github.Repository
	err  error
}

func (sc *spyChecker) CheckRate() ([]*github.Repository, error) {
	return sc.repo, sc.err
}
