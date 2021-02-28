package main

import (
	"context"
	"fmt"
	"os"

	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

// represents User object.
type user struct {
	Name           githubv4.String
	CommitComments struct {
		Nodes []commitComment
	} `graphql:"commitComments(last: 5)"`
	StarredRepositories struct {
		Nodes []repository
	} `graphql:"starredRepositories(last: 5)"`
}

// represents CommitComment object.
type commitComment struct {
	URL         githubv4.URI
	PublishedAt githubv4.DateTime
	CreatedAt   githubv4.DateTime
	UpdatedAt   githubv4.DateTime
}

// represents Repository object.
type repository struct {
	URL githubv4.URI
}

func main() {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	httpClient := oauth2.NewClient(context.Background(), src)

	client := githubv4.NewClient(httpClient)

	var query struct {
		Viewer struct {
			Following struct {
				Nodes []user
			} `graphql:"following(last: 5)"`
		}
	}

	err := client.Query(context.Background(), &query, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, v := range query.Viewer.Following.Nodes {
		fmt.Printf("name: %s, commits: %+v, starred: %+v\n", v.Name, v.CommitComments, v.StarredRepositories)
	}
}
