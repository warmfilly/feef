package main

import (
	"github.com/rivo/tview"
)

func main() {
	table := getTable()

	if err := tview.NewApplication().SetRoot(table, true).Run(); err != nil {
		panic(err)
	}
}

func getTable() *tview.Table {
	table := tview.NewTable().
		SetBorders(true).
		SetCell(0, 0, tview.NewTableCell("Cell One")).
		SetCell(0, 1, tview.NewTableCell("Cell Two")).
		SetCell(1, 0, tview.NewTableCell("Cell Three")).
		SetCell(1, 1, tview.NewTableCell("Cell Four"))

	return table
}
