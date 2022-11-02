package view

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

func PresentationUpdate(m Model, msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.Type == tea.KeyEsc || msg.Type == tea.KeyEscape || msg.String() == "q" {
			m.CurrentView = PRESENTATION_LIST

			return m, nil
		}
	}

	return m, nil
}

func PresentationView(m Model) string {
	presentation := m.SelectedPresentation

	firstLine := fmt.Sprintf("%s by %s", presentation.PresentationTitle, presentation.Presenter)
	secondLine := presentation.Desc
	thirdLine := strings.Repeat("⭐", presentation.Rating)
	fourthLine := presentation.Review

	completeView := fmt.Sprintf("%s\n%s\n%s\n%s\n", firstLine, secondLine, thirdLine, fourthLine)

	return docStyle.Render(completeView)
}
