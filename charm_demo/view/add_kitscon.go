package view

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/doverstav/kitscon222/charm_demo/commands"
)

var inputStyle = lipgloss.NewStyle().Margin(1, 0)

func AddKitsconUpdate(m Model, msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.Type == tea.KeyTab {
			m.Input.Blur()
			m.TextArea.Focus()

			return m, nil
		} else if msg.Type == tea.KeyShiftTab {
			m.TextArea.Blur()
			m.Input.Focus()

			return m, nil
		} else if msg.String() == "ctrl+j" { // Ctrl + Enter reads as ctrl+j for some reason
			return m, commands.SaveKitscon(m.DB, m.Input.Value(), m.TextArea.Value())
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
