package view

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

func PresentationListUpdate(m Model, msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "a" {
			// Add new presentation
			fmt.Print("Add presentation")
			m.CurrentView = ADD_NEW_PRESENTATION
			m.PresentationTitleInput.Focus()

			return m, nil
		} else if msg.Type == tea.KeyEnter {
			// Select presentation to view details
			fmt.Print("Selected presentation")
		}
	}

	var cmd tea.Cmd
	m.ItemList, cmd = m.ItemList.Update(msg)

	return m, cmd
}

func PresentationListView(m Model) string {
	return docStyle.Render(m.ItemList.View())
}
