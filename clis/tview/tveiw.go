package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/rivo/tview"
)

var ctype = flag.String("c", "test", "represents tview component type")

type ComponentType string

const (
	boxType   ComponentType = "box"
	listType  ComponentType = "list"
	formType  ComponentType = "form"
	pagesType ComponentType = "pages"
	modalType ComponentType = "modal"
)

func castComponent(s string) ComponentType {
	return ComponentType(s)
}

func usage() {
	fmt.Fprintf(os.Stderr, "usage: xxx [flags] [path ...]\n")
	flag.PrintDefaults()
	os.Exit(2)
}

// implenment as cli below demos.
// https://github.com/rivo/tview/wiki
// nolint:gocyclo
func main() {
	flag.Parse()
	flag.Usage = usage

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
	case formType:
		form := tview.NewForm().
			AddDropDown("Title", []string{"Mr.", "Ms.", "Mrs.", "Dr.", "Prof."}, 0, nil).
			AddInputField("First name", "", 20, nil, nil).
			AddInputField("Last name", "", 20, nil, nil).
			AddCheckbox("Age 18+", false, nil).
			AddPasswordField("Password", "", 10, '*', nil).
			AddButton("Save", nil).
			AddButton("Quit", func() {
				app.Stop()
			})
		form.SetBorder(true).SetTitle("Enter some data").SetTitleAlign(tview.AlignLeft)
		if err := app.SetRoot(form, true).SetFocus(form).Run(); err != nil {
			log.Fatal(err)
		}
	case pagesType:
		pageCount := 5
		pages := tview.NewPages()
		for page := 0; page < pageCount; page++ {
			func(page int) {
				pages.AddPage(fmt.Sprintf("page-%d", page),
					tview.NewModal().
						SetText(fmt.Sprintf("This is page %d. Choose where to go next.", page+1)).
						AddButtons([]string{"Next", "Quit"}).
						SetDoneFunc(func(buttonIndex int, buttonLabel string) {
							if buttonIndex == 0 {
								pages.SwitchToPage(fmt.Sprintf("page-%d", (page+1)%pageCount))
							} else {
								app.Stop()
							}
						}),
					false,
					page == 0)
			}(page)
		}
		if err := app.SetRoot(pages, true).SetFocus(pages).Run(); err != nil {
			log.Fatal(err)
		}
	case modalType:
		modal := tview.NewModal().
			SetText("Do you want to quit the application?").
			AddButtons([]string{"Quit", "Cancel"}).
			SetDoneFunc(func(buttonIndex int, buttonLabel string) {
				if buttonLabel == "Quit" {
					app.Stop()
				}
			})
		if err := app.SetRoot(modal, false).SetFocus(modal).Run(); err != nil {
			log.Fatal(err)
		}
	default:
		log.Printf("given component type is invalid: %s", ct)
		flag.Usage()
	}
}
