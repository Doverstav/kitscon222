package database

import (
	"encoding/json"
	"fmt"

	"github.com/dgraph-io/badger/v3"
	"github.com/google/uuid"
)

type ErrNoSuchPresentation struct{}

func (e *ErrNoSuchPresentation) Error() string {
	return "No such presentation"
}

type Presentation struct {
	Id                uuid.UUID `json:"id"`
	PresentationTitle string    `json:"title"`
	Presenter         string    `json:"presenter"`
	Desc              string    `json:"description"`
	Rating            int       `json:"rating"`
	Review            string    `json:"review"`
}

func GetPresentations(db *badger.DB, kitsconId uuid.UUID) ([]Presentation, error) {
	var kitscon Kitscon
	presentations := []Presentation{}
	err := GetItem(db, kitsconId.String(), &kitscon)
	if err != nil {
		fmt.Printf("Error getting kitscon %s: %v", kitsconId.String(), err)
		return []Presentation{}, err
	}

	for _, id := range kitscon.PresentationIds {
		var tempPres Presentation
		err := GetItem(db, id.String(), &tempPres)
		if err != nil {
			fmt.Printf("Error getting presentations %s: %v\n", id.String(), err)
		}

		presentations = append(presentations, tempPres)
	}

	return presentations, nil
}

func GetPresentationByName(db *badger.DB, kitsconId uuid.UUID, presentationName string) (Presentation, error) {
	var kitscon Kitscon
	err := GetItem(db, kitsconId.String(), &kitscon)
	if err != nil {
		return Presentation{}, err
	}

	for _, presId := range kitscon.PresentationIds {
		var tempPres Presentation
		GetItem(db, presId.String(), &tempPres)
		if tempPres.PresentationTitle == presentationName {
			return tempPres, nil
		}
	}

	return Presentation{}, &ErrNoSuchPresentation{}
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
