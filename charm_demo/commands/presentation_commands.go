package commands

import tea "github.com/charmbracelet/bubbletea"

func GetPresentations(kitsconId string) tea.Cmd {
	return func() tea.Msg {
		presentationList := []Presentation{
			{Id: "1", Title: "Hello", Presenter: "World", Rating: 4, Review: "Short and sweet"},
			{Id: "2", Title: "Goodbye", Presenter: "World", Rating: 2, Review: "Long and dull"},
		}

		return PresentationsMsg(presentationList)
	}
}

type Presentation struct {
	Id        string
	Title     string
	Presenter string
	Rating    int
	Review    string
}

type PresentationsMsg []Presentation
