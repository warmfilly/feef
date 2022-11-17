package main

import (
	"github.com/76creates/stickers"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	table *stickers.TableSingleType[string]
}

func (m Model) Init() tea.Cmd {
	return nil
}

// This is where all the keypresses will be handled for the app
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "up", "k":
			m.table.CursorUp()

		case "down", "j":
			m.table.CursorDown()
		}
	}

	return m, nil
}

func (m Model) View() string {
	return m.table.Render()
}
