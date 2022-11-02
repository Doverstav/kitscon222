package view

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/doverstav/kitscon222/charm_demo/commands"
)

var inputStyle = lipgloss.NewStyle().Margin(1, 0, 0, 0)

func AddKitsconUpdate(m Model, msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.Type == tea.KeyTab {
			m.KitsconTitleInput.Blur()
			m.KitsconDescriptionInput.Focus()

			return m, nil
		} else if msg.Type == tea.KeyShiftTab {
			m.KitsconDescriptionInput.Blur()
			m.KitsconTitleInput.Focus()

			return m, nil
		} else if msg.String() == "ctrl+j" { // Ctrl + Enter reads as ctrl+j for some reason
			return m, commands.SaveKitscon(m.DB, m.KitsconTitleInput.Value(), m.KitsconDescriptionInput.Value())
		}
	}

	var fieldCmd tea.Cmd
	m.KitsconTitleInput, fieldCmd = m.KitsconTitleInput.Update(msg)

	var areaCmd tea.Cmd
	m.KitsconDescriptionInput, areaCmd = m.KitsconDescriptionInput.Update(msg)

	return m, tea.Batch(fieldCmd, areaCmd)
}

func AddKitsconView(m Model) string {
	return docStyle.Render(
		inputStyle.Render(m.KitsconTitleInput.View()) + "\n" +
			inputStyle.Render(m.KitsconDescriptionInput.View()),
	)
}
