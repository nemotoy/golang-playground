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
	Name string
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

	table := tview.NewTable().
		SetBorders(false)
	rows := len(sources)
	for r := 0; r < rows; r++ {
		// func setCell([]struct){}
		data := sources[r]
		table.SetCell(r, 0,
			tview.NewTableCell(fmt.Sprint(data.ID)).
				SetAlign(tview.AlignLeft))
		table.SetCell(r, 1,
			tview.NewTableCell(fmt.Sprint(data.Name)).
				SetAlign(tview.AlignLeft))
		table.SetCell(r, 2,
			tview.NewTableCell(fmt.Sprint(data.URL)).
				SetAlign(tview.AlignLeft))
	}

	actionItems := tview.NewList().
		AddItem("Focus", "Press to focus", 'f', func() {
			app.SetFocus(table)
		}).
		AddItem("Quit", "Press to exit", 'q', func() {
			app.Stop()
		})

	grid := tview.NewGrid().
		SetRows(3, 0, 3).
		SetColumns(30, 0, 30).
		SetBorders(true).
		AddItem(table, 1, 0, 1, 3, 0, 100, false).
		AddItem(actionItems, 2, 0, 1, 3, 0, 100, false)

	app.SetRoot(grid, true).SetFocus(actionItems)

	if err := app.Run(); err != nil {
		panic(err)
	}
}
