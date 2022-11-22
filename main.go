package main

import (
	"github.com/rivo/tview"
	fsd "github.com/warmfilly/feef/filesystemDisplay"
)

func main() {
	table := fsd.GetTable("/home/warm/")

	if err := tview.NewApplication().SetRoot(table, true).Run(); err != nil {
		panic(err)
	}
}
