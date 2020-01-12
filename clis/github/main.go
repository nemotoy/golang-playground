package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"os"
	"strings"

	"strconv"

	"github.com/google/go-github/v28/github"
	"github.com/olekukonko/tablewriter"
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

		// procces users
		row := make([][]string, len(users))
		for i, u := range users {
			id := strconv.Itoa((i))
			row[i] = []string{id, *u.HTMLURL}
		}

		// output a table
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"ID", "URL"})
		for _, r := range row {
			table.Append(r)
		}
		table.Render()

		// input stdin
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Input ID: ")
		s, _ := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintf(os.Stderr, "%#v\n", err)
			break
		}

		// output stdout
		s = strings.Trim(s, " \n")
		id, err := strconv.Atoi(s)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%#v\n", err)
			break
		}
		for i, u := range users {
			if i == id {
				fmt.Println(u)
			}
		}
		// TODO: search activity from the selected user
	default:
		fmt.Fprintf(os.Stderr, "Invalid type: %s\n", *apiType)
		exitCode = 2
		return
	}
}
