package view

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var inputStyle = lipgloss.NewStyle().Margin(1, 0)

func AddKitsconUpdate(m Model, msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		fmt.Print("test")
		if msg.Type == tea.KeyTab {
			m.Input.Blur()
			m.TextArea.Focus()

			return m, nil
		}
	}

	var fieldCmd tea.Cmd
	m.Input, fieldCmd = m.Input.Update(msg)

	var areaCmd tea.Cmd
	m.TextArea, areaCmd = m.TextArea.Update(msg)

	return m, tea.Batch(fieldCmd, areaCmd)
}

func AddKitsconView(m Model) string {
	return docStyle.Render(inputStyle.Render(m.Input.View()) + "\n" + m.TextArea.View())
}
