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
	flexType  ComponentType = "flex"
	gridType  ComponentType = "grid"
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
	case flexType:
		flex := tview.NewFlex().
			AddItem(tview.NewBox().SetBorder(true).SetTitle("Left (1/2 x width of Top)"), 0, 1, false).
			AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
				AddItem(tview.NewBox().SetBorder(true).SetTitle("Top"), 0, 1, false).
				AddItem(tview.NewBox().SetBorder(true).SetTitle("Middle (3 x height of Top)"), 0, 3, false).
				AddItem(tview.NewBox().SetBorder(true).SetTitle("Bottom (5 rows)"), 5, 1, false), 0, 2, false).
			AddItem(tview.NewBox().SetBorder(true).SetTitle("Right (20 cols)"), 20, 1, false)
		if err := app.SetRoot(flex, true).SetFocus(flex).Run(); err != nil {
			panic(err)
		}
	case gridType:
		newPrimitive := func(text string) tview.Primitive {
			return tview.NewTextView().
				SetTextAlign(tview.AlignCenter).
				SetText(text)
		}
		menu := newPrimitive("Menu")
		// main := newPrimitive("Main content")
		sideBar := newPrimitive("Side Bar")
		listItems := tview.NewList().
			AddItem("List item 1", "Some explanatory text", 'a', nil).
			AddItem("List item 2", "Some explanatory text", 'b', nil).
			AddItem("List item 3", "Some explanatory text", 'c', nil).
			AddItem("List item 4", "Some explanatory text", 'd', nil).
			AddItem("Quit", "Press to exit", 'q', func() {
				app.Stop()
			})

		grid := tview.NewGrid().
			SetRows(3, 0, 3).
			SetColumns(30, 0, 30).
			SetBorders(true).
			AddItem(newPrimitive("Header"), 0, 0, 1, 3, 0, 0, false).
			AddItem(newPrimitive("Footer"), 2, 0, 1, 3, 0, 0, false)

		// // Layout for screens narrower than 100 cells (menu and side bar are hidden).
		// grid.AddItem(menu, 0, 0, 0, 0, 0, 0, false).
		// 	AddItem(main, 1, 0, 1, 3, 0, 0, false).
		// 	AddItem(sideBar, 0, 0, 0, 0, 0, 0, false)

		// Layout for screens wider than 100 cells.
		grid.AddItem(menu, 1, 0, 1, 1, 0, 100, false).
			AddItem(listItems, 1, 1, 1, 1, 0, 100, false).
			AddItem(sideBar, 1, 2, 1, 1, 0, 100, false)

		app.SetRoot(grid, true)
		app.SetFocus(listItems)

		if err := app.Run(); err != nil {
			panic(err)
		}
	default:
		log.Printf("given component type is invalid: %s", ct)
		flag.Usage()
	}
}
