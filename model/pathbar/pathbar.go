package pathbar

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type Pathbar struct {
	Path string
}

func (pathbar *Pathbar) Init() tea.Cmd {
	return nil
}

func (pathbar *Pathbar) Update(msg tea.Msg) (Pathbar, tea.Cmd) {
	return *pathbar, nil
}

func (pathbar *Pathbar) View() string {
	sb := strings.Builder{}

	pathParts := strings.Split(pathbar.Path, "/")

	for i, part := range pathParts {
		if part == "" {
			continue
		}

		sb.WriteString(part)

		if i <= len(pathParts) {
			sb.WriteString(" -> ")
		}
	}

	return sb.String()
}
