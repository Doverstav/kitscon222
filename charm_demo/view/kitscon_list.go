package view

import (
	"github.com/charmbracelet/lipgloss"
)

var listStyle = lipgloss.NewStyle().Margin(1, 2)

func KitsConListUpdate() {

}

func KitsConListView(m Model) string {
	return listStyle.Render(m.List.View())
}
