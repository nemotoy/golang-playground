package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gdamore/tcell/v2"
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
	// flag
	flag.Parse()
	flag.Usage = usage

	// fetch data
	sources := []source{
		{1, "Go", "https://golang.org/"},
		{2, "Rust", "https://www.rust-lang.org/"},
		{3, "Kotlin", "https://kotlinlang.org/"},
	}

	// new tview app
	app := tview.NewApplication()

	table := tview.NewTable().
		SetBorders(false)
	table.SetDoneFunc(func(key tcell.Key) {
		switch key {
		case tcell.KeyEscape:
			app.Stop()
		case tcell.KeyEnter:
			table.SetSelectable(true, true)
			// restrict selectable clumns and add donefunc
		}
	})
	// new layout
	rows := len(sources)
	for r := 0; r < rows; r++ {
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

	// new layout
	grid := tview.NewGrid().
		SetRows(3, 0, 3).
		SetColumns(30, 0, 30).
		SetBorders(true).
		AddItem(table, 1, 0, 1, 3, 0, 100, false).
		AddItem(actionItems, 2, 0, 1, 3, 0, 100, false)

	if err := app.SetRoot(grid, true).SetFocus(actionItems).Run(); err != nil {
		panic(err)
	}
}
