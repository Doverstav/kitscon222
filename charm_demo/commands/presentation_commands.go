package commands

import tea "github.com/charmbracelet/bubbletea"

func GetPresentations(kitsconId string) tea.Cmd {
	return func() tea.Msg {
		presentationList := []Presentation{
			{id: "1", title: "Hello", presenter: "World", rating: 4, review: "Short and sweet"},
			{id: "2", title: "Goodbye", presenter: "World", rating: 2, review: "Long and dull"},
		}

		return PresentationsMsg(presentationList)
	}
}

type Presentation struct {
	id        string
	title     string
	presenter string
	rating    int
	review    string
}

type PresentationsMsg []Presentation
