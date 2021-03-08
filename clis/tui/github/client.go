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
	// ContributionsCollection struct {
	// 	CommitContributionsByRepository []commitContributionsByRepository
	// 	TotalCommitContributions        githubv4.Int
	// 	TotalIssueContributions         githubv4.Int
	// 	TotalPullRequestContributions   githubv4.Int
	// } `graphql:"contributionsCollection(from: $contributionsCollectionFrom, to: $contributionsCollectionTo)"`
}

// represents StarredRepositoryEdge object.
type starredRepositoritoryEdge struct {
	StarredAt githubv4.DateTime
}

// represents Repository object.
type repository struct {
	URL githubv4.URI
	// https://docs.github.com/en/graphql/reference/input-objects#languageorder
	Languages struct {
		Nodes []language
	} `graphql:"languages(last:10)"`
	PrimaryLanguage struct {
		Color githubv4.String
		Name  githubv4.String
	}
	StargazerCount githubv4.Int
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
		CommitUrl     githubv4.URI
		CommittedDate githubv4.DateTime
	}
}

func main() {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	httpClient := oauth2.NewClient(context.Background(), src)

	client := githubv4.NewClient(httpClient)

	// user activities
	// - commit
	//   - language
	//   - commit url
	//   - repo name
	//   - created_at
	// - starred
	// - following
	var query struct {
		Viewer struct {
			Following struct {
				Nodes []user
			} `graphql:"following(last: $followingLast)"`
		}
	}

	var from, to struct{ time.Time }
	now := time.Now()
	from.Time = now.Add(-24 * time.Hour)
	to.Time = now
	variables := map[string]interface{}{
		"followingLast":           githubv4.Int(5),
		"starredRepositoriesLast": githubv4.Int(5),
		// "contributionsCollectionFrom": githubv4.DateTime(from),
		// "contributionsCollectionTo":   githubv4.DateTime(to),
	}

	err := client.Query(context.Background(), &query, variables)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, v := range query.Viewer.Following.Nodes {
		fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++++++++++")
		fmt.Printf("name: %s\n", v.Name)
		fmt.Println("Starred repositories: ")
		for i, st := range v.StarredRepositories.Nodes {
			// length of edges equal nodes
			fmt.Printf("StarredAt: %v, Languages: %s, , Count: %d, URL: %s\n", v.StarredRepositories.Edges[i].StarredAt, st.PrimaryLanguage.Name, st.StargazerCount, st.URL)
		}
		// fmt.Println("- Commits")
		// for _, ccr := range v.ContributionsCollection.CommitContributionsByRepository {
		// 	for _, cn := range ccr.Contributions.Nodes {
		// 		for _, p := range cn.Repository.PullRequests.Nodes {
		// 			fmt.Printf("language: %+v, commits: %+v\n", cn.Repository.Languages, p.Commits.Nodes)
		// 		}
		// 	}
		// }
		fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++++++++++")
	}
}
