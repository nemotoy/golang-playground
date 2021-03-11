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
	ID       int
	Language string
	URL      string
}

// nolint:gocyclo
func main() {
	// parse flags
	flag.Parse()
	flag.Usage = usage

	// fetch data
	sources := []source{}
	for i := 1; i <= 10; i++ {
		sources = append(sources, source{ID: i, Language: "lang", URL: "https://example.org"})
	}

	// new tview app
	app := tview.NewApplication()

	// new table
	table := tview.NewTable().
		SetBorders(false).
		SetFixed(1, 1)
	table.SetDoneFunc(func(key tcell.Key) {
		switch key {
		case tcell.KeyEscape:
			app.Stop()
		case tcell.KeyEnter:
			table.SetSelectable(true, true)
		}
	})
	// add frame(row 0)
	table.SetCell(0, 1,
		tview.NewTableCell("ID").
			SetAlign(tview.AlignLeft))
	table.SetCell(0, 2,
		tview.NewTableCell("Language").
			SetAlign(tview.AlignLeft))
	table.SetCell(0, 3,
		tview.NewTableCell("URL").
			SetAlign(tview.AlignLeft))
	// add new layout
	rows := len(sources)
	for r := 1; r <= rows; r++ {
		data := sources[r-1]
		table.SetCell(r, 1,
			tview.NewTableCell(fmt.Sprint(data.ID)).
				SetAlign(tview.AlignLeft))
		table.SetCell(r, 2,
			tview.NewTableCell(fmt.Sprint(data.Language)).
				SetAlign(tview.AlignLeft))
		table.SetCell(r, 3,
			tview.NewTableCell(fmt.Sprint(data.URL)).
				SetAlign(tview.AlignLeft))
	}
	// new action fields
	actionItems := tview.NewList().
		AddItem("Scroll beginning table", "", 'b', func() {
			table.ScrollToBeginning()
		}).
		AddItem("Scroll end table", "", 'e', func() {
			table.ScrollToEnd()
		}).
		AddItem("Focus", "", 'f', func() {
			app.SetFocus(table)
		}).
		AddItem("Quit", "", 'q', func() {
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
