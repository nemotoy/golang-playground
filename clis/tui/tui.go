package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/rivo/tview"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: xxx [flags] [path ...]\n")
	flag.PrintDefaults()
	os.Exit(2)
}

type source struct {
	ID   int
	name string
	URL  string
}

// nolint:gocyclo
func main() {
	flag.Parse()
	flag.Usage = usage

	app := tview.NewApplication()

	sources := []source{
		{1, "Go", "https://golang.org/"},
		{2, "Rust", "https://www.rust-lang.org/"},
		{3, "Kotlin", "https://kotlinlang.org/"},
	}

	listItems := tview.NewList()
	for _, v := range sources {
		listItems.AddItem(v.name, v.URL, '0', nil)
	}
	listItems.AddItem("Quit", "Press to exit", 'q', func() {
		app.Stop()
	})

	grid := tview.NewGrid().
		SetRows(3, 0, 3).
		SetColumns(30, 0, 30).
		SetBorders(true).
		AddItem(listItems, 1, 1, 1, 1, 0, 100, false)

	app.SetRoot(grid, true).SetFocus(listItems)

	if err := app.Run(); err != nil {
		panic(err)
	}
}
