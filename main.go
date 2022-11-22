package main

import (
	"io/fs"
	"io/ioutil"
	"log"

	"github.com/rivo/tview"
)

func main() {
	table := getTable("/home/warm/")

	if err := tview.NewApplication().SetRoot(table, true).Run(); err != nil {
		panic(err)
	}
}

func getTable(path string) *tview.Table {
	table := tview.NewTable().
		SetBorders(true)

	table.SetFixed(1, 0)

	table.SetCell(0, 0, tview.NewTableCell("Name"))
	table.SetCell(0, 1, tview.NewTableCell("Size"))

	for index, file := range iterateFiles(path) {
		table.SetCell(index+1, 0, tview.NewTableCell(file.Name()))
		table.SetCell(index+1, 1, tview.NewTableCell("10kb"))
	}

	return table
}

func iterateFiles(path string) []fs.FileInfo {
	files, err := ioutil.ReadDir(path)

	if err != nil {
		log.Fatal(err)
	}

	return files
}
