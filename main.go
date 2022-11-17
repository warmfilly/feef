package main

import (
	"fmt"
	"os"

	"github.com/76creates/stickers"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	headers := []string{"File"}
	m := Model{
		table: stickers.NewTableSingleType[string](50, 10, headers),
	}

	ratio := []int{1}
	//minSize := []int{4, 5, 5, 2, 5}
	m.table.SetRatio(ratio)

	rows := getRows()
	m.table.AddRows(rows)

	p := tea.NewProgram(m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}

func getRows() [][]string {
	var matrix [][]string
	row1 := []string{"one"}
	row2 := []string{"two"}

	matrix = append(matrix, row1)
	matrix = append(matrix, row2)

	return matrix
}
