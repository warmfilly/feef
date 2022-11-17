package statusbar

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Statusbar struct {
	Status string
}

func (statusbar *Statusbar) Init() tea.Cmd {
	return nil
}

func (statusbar *Statusbar) Update(msg tea.Msg) (Statusbar, tea.Cmd) {
	return *statusbar, nil
}

func (statusbar *Statusbar) View() string {
	return statusbar.Status
}
