package view

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/doverstav/kitscon222/charm_demo/commands"
)

func PresentationListUpdate(m Model, msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "a" {
			// Add new presentation
			m.CurrentView = ADD_NEW_PRESENTATION
			m.PresentationTitleInput.Focus()

			return m, nil
		} else if msg.String() == "d" {
			selectedItem, ok := m.ItemList.SelectedItem().(commands.Presentation)
			if !ok {
				fmt.Printf("Could not convert %v to Presentation", m.ItemList.SelectedItem())
				return m, nil
			}

			return m, commands.RemovePresentation(m.DB, m.SelectedKitscon.Id, selectedItem.Id)
		} else if msg.Type == tea.KeyEnter {
			// Select presentation to view details
			fmt.Print("Selected presentation")
			m.CurrentView = PRESENTATION
			m.SelectedPresentation, _ = m.ItemList.SelectedItem().(commands.Presentation)

			return m, nil
		}
	}

	var cmd tea.Cmd
	m.ItemList, cmd = m.ItemList.Update(msg)

	return m, cmd
}

func PresentationListView(m Model) string {
	return docStyle.Render(m.ItemList.View())
}
