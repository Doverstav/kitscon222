package view

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/doverstav/kitscon222/charm_demo/commands"
)

func KitsConListUpdate(m Model, msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "a" {
			m.CurrentView = ADD_NEW_KITSCON
			m.KitsconTitleInput.Focus()
			return m, nil
		} else if msg.Type == tea.KeyEnter {
			selectedKitscon, _ := m.List.SelectedItem().(commands.Kitscon)
			return m, commands.KitsconSelected(selectedKitscon)
		}
	}
	var cmd tea.Cmd
	m.List, cmd = m.List.Update(msg)

	return m, cmd
}

func KitsConListView(m Model) string {
	return docStyle.Render(m.List.View())
}
