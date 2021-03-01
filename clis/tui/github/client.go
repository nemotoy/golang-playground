package main

import (
	"context"
	"fmt"
	"os"
	"time"

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
	TopRepositories struct {
		Nodes []repository
	} `graphql:"topRepositories(last: $topRepositoriesLast, orderBy: {field: CREATED_AT, direction: ASC})"`
	ContributionsCollection struct {
		CommitContributionsByRepository []commitContributionsByRepository
		TotalCommitContributions        githubv4.Int
		TotalIssueContributions         githubv4.Int
		TotalPullRequestContributions   githubv4.Int
	} `graphql:"contributionsCollection(from: $contributionsCollectionFrom, to: $contributionsCollectionTo)"`
}

// represents StarredRepositoryEdge object.
type starredRepositoritoryEdge struct {
	StarredAt githubv4.DateTime
}

// represents Repository object.
type repository struct {
	URL githubv4.URI
}

type commitContributionsByRepository struct {
	Contributions struct {
		Nodes      []createdCommitContribution
		TotalCount githubv4.Int
	} `graphql:"contributions(last:10)"`
}

type createdCommitContribution struct {
	CommitCount githubv4.Int
	OccurredAt  githubv4.DateTime
	URL         githubv4.URI
	Repository  repositoryPerCommit
}

type repositoryPerCommit struct {
	Languages struct {
		Nodes []language
	} `graphql:"languages(last:10)"`
	PullRequests struct {
		Nodes []pullRequest
	} `graphql:"pullRequests(last:10)"`
}

type language struct {
	Color githubv4.String
	Name  githubv4.String
}

type pullRequest struct {
	Commits struct {
		Nodes []pullRequestCommit
	} `graphql:"commits(last:10)"`
}

type pullRequestCommit struct {
	Commit struct {
		CommitUrl githubv4.URI
	}
}

func main() {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	httpClient := oauth2.NewClient(context.Background(), src)

	client := githubv4.NewClient(httpClient)

	// todo:
	// - commit
	// - starred
	// - following
	var query struct {
		Viewer struct {
			Following struct {
				Nodes []user
			} `graphql:"following(last: $followingLast)"`
		}
	}

	now := time.Now()
	variables := map[string]interface{}{
		"followingLast":               githubv4.Int(5),
		"starredRepositoriesLast":     githubv4.Int(5),
		"topRepositoriesLast":         githubv4.Int(5),
		"contributionsCollectionFrom": githubv4.DateTime{now.Add(-24 * time.Hour)},
		"contributionsCollectionTo":   githubv4.DateTime{now},
	}

	err := client.Query(context.Background(), &query, variables)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, v := range query.Viewer.Following.Nodes {
		fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++++++++++")
		fmt.Printf("name: %s, edges:%+v\n", v.Name, v.StarredRepositories.Edges)
		for _, st := range v.StarredRepositories.Nodes {
			fmt.Printf("URL: %s\n", st.URL)
		}
		for _, tr := range v.TopRepositories.Nodes {
			fmt.Printf("URL: %+v\n", tr)
		}
		fmt.Printf("Commit contribution: %+v\n", v.ContributionsCollection.CommitContributionsByRepository)
		fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++++++++++")
	}
}
