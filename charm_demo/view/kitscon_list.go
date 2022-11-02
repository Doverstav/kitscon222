package view

import (
	tea "github.com/charmbracelet/bubbletea"
)

func KitsConListUpdate(m Model, msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "a" {
			m.CurrentView = ADD_NEW_KITSCON
			m.Input.Focus()
			return m, nil
		}
	}
	var cmd tea.Cmd
	m.List, cmd = m.List.Update(msg)

	return m, cmd
}

func KitsConListView(m Model) string {
	return docStyle.Render(m.List.View())
}
