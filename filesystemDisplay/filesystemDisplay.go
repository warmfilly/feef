package filesystemDisplay

import (
	"io/fs"
	"io/ioutil"
	"log"

	"github.com/dustin/go-humanize"
	"github.com/rivo/tview"
)

func GetTable(path string) *tview.Table {
	table := initTable()

	for index, file := range iterateFiles(path) {
		var nameCell *tview.TableCell

		if file.IsDir() {
			nameCell = dirCell(file.Name())
		} else {
			nameCell = fileCell(file.Name())
		}
		table.SetCell(index+1, 0, nameCell)

		sizeStr := humanize.Bytes(uint64(file.Size()))
		table.SetCell(index+1, 1, tview.NewTableCell(sizeStr))

	}

	return table
}

func fileCell(name string) *tview.TableCell {
	cell := tview.NewTableCell("[red]" + name)
	return cell
}

func dirCell(name string) *tview.TableCell {
	cell := tview.NewTableCell("[blue]" + name)
	return cell
}

func initTable() *tview.Table {
	table := tview.NewTable().
		SetBorders(true)

	table.SetFixed(1, 0)

	table.SetCell(0, 0, tview.NewTableCell("Name"))
	table.SetCell(0, 1, tview.NewTableCell("Size"))

	return table
}

func iterateFiles(path string) []fs.FileInfo {
	files, err := ioutil.ReadDir(path)

	if err != nil {
		log.Fatal(err)
	}

	return files
}
