package database

import (
	"encoding/json"
	"fmt"

	"github.com/dgraph-io/badger/v3"
	"github.com/google/uuid"
)

type Presentation struct {
	Id                uuid.UUID `json:"id"`
	PresentationTitle string    `json:"title"`
	Presenter         string    `json:"presenter"`
	Desc              string    `json:"description"`
	Rating            int       `json:"rating"`
	Review            string    `json:"review"`
}

func GetPresentations(db *badger.DB, kitsconName string) []Presentation {
	kitscons := GetKitscons(db)

	presentations := []Presentation{}
	for _, kitscon := range kitscons {
		// TODO This will amthc several times in case of duped names. Is this a problem?
		if kitscon.Name == kitsconName {
			for _, id := range kitscon.PresentationIds {
				var tempPres Presentation
				GetItem(db, id.String(), &tempPres)
				presentations = append(presentations, tempPres)
			}
		}
	}

	return presentations
}

func SavePresentation(db *badger.DB, kitsconId uuid.UUID, title string, presenter string, description string, rating int, review string) error {
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
		return err
	}

	// Save to database
	err = SaveItem(db, newPresentation.Id.String(), marshalled)
	if err != nil {
		fmt.Printf("Error when saving %v: %v", newPresentation, err)
		return err
	}

	// Get parent KitsCon, add presentation to it
	var kitscon Kitscon
	err = GetItem(db, kitsconId.String(), &kitscon)
	if err != nil {
		fmt.Printf("Failed to fetch kitscon %s: %v", kitsconId.String(), err)
	}
	kitscon.PresentationIds = append(kitscon.PresentationIds, newPresentation.Id)

	marshalled, err = json.Marshal(kitscon)
	if err != nil {
		fmt.Printf("Error when marshalling %v: %v", kitscon, err)
		return err
	}

	// Update KitsCon in database
	err = SaveItem(db, kitsconId.String(), marshalled)
	if err != nil {
		fmt.Printf("Error when saving %v: %v", kitscon, err)
		return err
	}

	return nil

}
