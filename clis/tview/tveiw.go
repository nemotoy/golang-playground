package main

import (
	"flag"
	"log"
	"os"

	"github.com/rivo/tview"
)

var ctype = flag.String("ct", "test", "represents tview component type")

type ComponentType string

const (
	boxType  ComponentType = "box"
	listType ComponentType = "list"
)

func castComponent(s string) ComponentType {
	return ComponentType(s)
}

// implenment as cli below demos.
// https://github.com/rivo/tview/wiki
func main() {
	flag.Parse()
	if *ctype == "" {
		os.Exit(1)
	}
	app := tview.NewApplication()
	ct := castComponent(*ctype)
	switch ct {
	case boxType:
		box := tview.NewBox().
			SetBorder(true).
			SetTitle("Box Demo")
		if err := app.SetRoot(box, true).Run(); err != nil {
			log.Fatal(err)
		}
	case listType:
		list := tview.NewList().
			AddItem("List item 1", "Some explanatory text", 'a', nil).
			AddItem("List item 2", "Some explanatory text", 'b', nil).
			AddItem("List item 3", "Some explanatory text", 'c', nil).
			AddItem("List item 4", "Some explanatory text", 'd', nil).
			AddItem("Quit", "Press to exit", 'q', func() {
				app.Stop()
			})
		if err := app.SetRoot(list, true).SetFocus(list).Run(); err != nil {
			log.Fatal(err)
		}
	default:
		log.Printf("given component type is not found: %s", ct)
	}
}
