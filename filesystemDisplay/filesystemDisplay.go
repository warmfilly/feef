package filesystemDisplay

import (
	"io/fs"
	"io/ioutil"
	"log"

	"github.com/dustin/go-humanize"
	"github.com/rivo/tview"
)

func GetTable(path string) *tview.Table {
	table := tview.NewTable().
		SetBorders(true)

	table.SetFixed(1, 0)

	table.SetCell(0, 0, tview.NewTableCell("Name"))
	table.SetCell(0, 1, tview.NewTableCell("Size"))

	for index, file := range iterateFiles(path) {
		table.SetCell(index+1, 0, tview.NewTableCell(file.Name()))

		sizeInt := uint64(file.Size())
		sizeStr := humanize.Bytes(sizeInt)
		table.SetCell(index+1, 1, tview.NewTableCell(sizeStr))
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
