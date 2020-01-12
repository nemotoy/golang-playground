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
	userName  = flag.String("u", "", "user name")
	apiType   = flag.String("t", "", "use this API type")
	repoLists = "list"
	followers = "following"
)

func main() {
	doMain()
	os.Exit(exitCode)
}

func doMain() {
	flag.Parse()
	client := github.NewClient(nil)
	switch *apiType {
	case repoLists:
		opt := &github.RepositoryListOptions{Type: "public"}
		if *userName == "" {
			fmt.Fprintf(os.Stderr, "the given user name is empty\n")
			exitCode = 2
			return
		}
		repos, resp, err := client.Repositories.List(context.Background(), *userName, opt)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to get list of repositories: %#v\n", err)
			exitCode = 2
			return
		}
		defer resp.Body.Close()
		fmt.Printf("First repo: %v\n, Rate: %s\n", repos[0], resp.Rate.String())
	case followers:
		users, resp, err := client.Users.ListFollowing(context.Background(), *userName, nil)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to get list of following: %#v\n", err)
			exitCode = 2
			return
		}
		defer resp.Body.Close()
		fmt.Printf("First following: %v\n, Rate: %s\n", users[0], resp.Rate.String())
	default:
		fmt.Fprintf(os.Stderr, "Invalid type: %s\n", *apiType)
		exitCode = 2
		return
	}
}
