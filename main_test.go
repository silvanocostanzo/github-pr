package main

import (
	"log"
	"reflect"
	"testing"

	"github.com/google/go-github/v41/github"
)

func Test_checkRateLimit(t *testing.T) {
	spyChecker := &SpyChecker{}

	got, err := checkRateLimit(spyChecker)
	want := make([]*github.Repository, 0)

	if err != nil {
		log.Fatal(err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got type %T, want type %T", got, want)
	}
}

type SpyChecker struct{}

func (sc *SpyChecker) CheckRate() ([]*github.Repository, error) {
	var repo = []*github.Repository{}
	return repo, nil
}
