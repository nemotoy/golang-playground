package main

import (
	"context"
	"fmt"
	"os"

	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

func main() {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	httpClient := oauth2.NewClient(context.Background(), src)

	client := githubv4.NewClient(httpClient)

	var query struct {
		Viewer struct {
			Name      githubv4.String
			Following struct {
				Nodes []struct {
					Name           githubv4.String
					CommitComments struct {
						Nodes []struct {
							URL         githubv4.URI
							PublishedAt githubv4.DateTime
							CreatedAt   githubv4.DateTime
							UpdatedAt   githubv4.DateTime
						}
					} `graphql:"commitComments(last: 5)"`
					StarredRepositories struct {
						Edges []struct {
							StarredAt githubv4.DateTime
						}
						Nodes []struct {
							URL githubv4.URI
						}
					} `graphql:"starredRepositories(last: 5)"`
				}
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
