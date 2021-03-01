# tui

## TODO

- main view
  - coordinate
- content view
- user actiion
  - key binding
  - change forcus each items(want apply action all views)
- data source
  - initial
  - refresh
- colors

## doc

### TUI

- https://pkg.go.dev/github.com/rivo/tview
  - https://github.com/xxjwxc/gormt
  - https://github.com/wtfutil/wtf
  - https://github.com/skanehira/docui
  - https://github.com/derailed/k9s

### GitHub Client

- v3 https://github.com/google/go-github/
- v4 https://github.com/shurcooL/githubv4/
  - https://docs.github.com/en/graphql/overview/explorer
  - https://docs.github.com/ja/graphql/guides/introduction-to-graphql

parent(repository) -> field(issue) -> field(comments) -> ...

```golang
var q struct {
	Repository struct {
		Issue struct {
			Comments struct {
				Nodes    []comment
				PageInfo struct {
					EndCursor   githubv4.String
					HasNextPage bool
				}
			} `graphql:"comments(first: 100, after: $commentsCursor)"` // 100 per page.
		} `graphql:"issue(number: $issueNumber)"`
	} `graphql:"repository(owner: $repositoryOwner, name: $repositoryName)"`
}
variables := map[string]interface{}{
	"repositoryOwner": githubv4.String(owner),
	"repositoryName":  githubv4.String(name),
	"issueNumber":     githubv4.Int(issue),
	"commentsCursor":  (*githubv4.String)(nil), // Null after argument to get first page.
}
```

ルートが親で子にオブジェクト。

- root: https://docs.github.com/en/graphql/reference/queries
- resources: https://docs.github.com/en/graphql/reference/objects

---

## GraphQL

- 深いネストの取り扱い
- 上から下への依存度
