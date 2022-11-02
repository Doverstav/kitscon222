package commands

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/dgraph-io/badger/v3"
	"github.com/doverstav/kitscon222/charm_demo/database"
	"github.com/google/uuid"
)

func GetPresentations(db *badger.DB, presentationIds []uuid.UUID) tea.Cmd {
	return func() tea.Msg {
		presentations := []Presentation{}

		for _, id := range presentationIds {
			var tempPres Presentation
			err := database.GetItem(db, id.String(), &tempPres)
			if err != nil {
				fmt.Printf("Error getting presentations %s: %v\n", id.String(), err)
			}

			presentations = append(presentations, tempPres)
		}

		// presentationList := []Presentation{
		// 	{Id: "1", Title: "Hello", Presenter: "World", Rating: 4, Review: "Short and sweet"},
		// 	{Id: "2", Title: "Goodbye", Presenter: "World", Rating: 2, Review: "Long and dull"},
		// }

		return PresentationsMsg(presentations)
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
