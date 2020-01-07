package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/google/go-github/v28/github"
)

var (
	exitCode = 0
)

var (
	userName = flag.String("u", "", "user name")
)

func main() {
	doMain()
	os.Exit(exitCode)
}

func doMain() {
	flag.Parse()
	client := github.NewClient(nil)
	opt := &github.RepositoryListOptions{Type: "public"}
	if *userName == "" {
		fmt.Fprintf(os.Stderr, "the given user name is empty\n")
		exitCode = 2
		return
	}
	repos, _, err := client.Repositories.List(context.Background(), *userName, opt)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to get list of repositories: %#v\n", err)
		exitCode = 2
		return
	}
	fmt.Println(repos)
}
