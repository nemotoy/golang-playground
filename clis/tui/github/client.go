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
	Name                githubv4.String
	StarredRepositories struct {
		Edges []starredRepositoritoryEdge
		Nodes []repository
	} `graphql:"starredRepositories(last: $starredRepositoriesLast)"`
}

// represents StarredRepositoryEdge object.
type starredRepositoritoryEdge struct {
	StarredAt githubv4.DateTime
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
			} `graphql:"following(last: $followingLast)"`
		}
	}

	variables := map[string]interface{}{
		"followingLast":           githubv4.Int(5),
		"starredRepositoriesLast": githubv4.Int(5),
	}

	err := client.Query(context.Background(), &query, variables)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, v := range query.Viewer.Following.Nodes {
		fmt.Printf("name: %s, edges:%+v\n", v.Name, v.StarredRepositories.Edges)
		for _, st := range v.StarredRepositories.Nodes {
			fmt.Printf("URL: %s\n", st.URL)
		}
	}
}
