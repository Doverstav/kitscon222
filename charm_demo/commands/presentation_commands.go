package commands

import (
	"encoding/json"
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/dgraph-io/badger/v3"
	"github.com/doverstav/kitscon222/charm_demo/database"
	"github.com/google/uuid"
)

func GetPresentations(db *badger.DB, kitsconId uuid.UUID) tea.Cmd {
	return func() tea.Msg {
		var kitscon Kitscon
		presentations := []Presentation{}
		err := database.GetItem(db, kitsconId.String(), &kitscon)
		if err != nil {
			fmt.Printf("Error getting kitscon %s: %v", kitsconId.String(), err)
			return PresentationsMsg(presentations)
		}

		for _, id := range kitscon.PresentationIds {
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

func SavePresentation(db *badger.DB, kitsconId uuid.UUID, title string, presenter string, description string, rating int, review string) tea.Cmd {
	return func() tea.Msg {
		// Create new presentation
		newPresentation := Presentation{
			Id:                uuid.New(),
			PresentationTitle: title,
			Presenter:         presenter,
			Desc:              description,
			Rating:            rating,
			Review:            review,
		}
		marshalled, err := json.Marshal(newPresentation)
		if err != nil {
			fmt.Printf("Error when marshalling %v: %v", newPresentation, err)
			return PresentationAddedMsg(false)
		}

		// Save to database
		err = database.SaveItem(db, newPresentation.Id.String(), marshalled)
		if err != nil {
			fmt.Printf("Error when saving %v: %v", newPresentation, err)
			return PresentationAddedMsg(false)
		}

		// Get parent KitsCon, add presentation to it
		var kitscon Kitscon
		err = database.GetItem(db, kitsconId.String(), &kitscon)
		if err != nil {
			fmt.Printf("Failed to fetch kitscon %s: %v", kitsconId.String(), err)
		}
		kitscon.PresentationIds = append(kitscon.PresentationIds, newPresentation.Id)

		marshalled, err = json.Marshal(kitscon)
		if err != nil {
			fmt.Printf("Error when marshalling %v: %v", kitscon, err)
			return PresentationAddedMsg(false)
		}

		// Update KitsCon in database
		err = database.SaveItem(db, kitsconId.String(), marshalled)
		if err != nil {
			fmt.Printf("Error when saving %v: %v", kitscon, err)
			return PresentationAddedMsg(false)
		}

		return PresentationAddedMsg(true)
	}
}

type Presentation struct {
	Id                uuid.UUID `json:"id"`
	PresentationTitle string    `json:"title"`
	Presenter         string    `json:"presenter"`
	Desc              string    `json:"description"`
	Rating            int       `json:"rating"`
	Review            string    `json:"review"`
}

func (p Presentation) Title() string {
	return fmt.Sprintf("%s by %s", p.PresentationTitle, p.Presenter)
}
func (p Presentation) Description() string { return strings.Repeat("⭐", p.Rating) }
func (p Presentation) FilterValue() string { return p.Title() }

type PresentationsMsg []Presentation

type PresentationAddedMsg bool
