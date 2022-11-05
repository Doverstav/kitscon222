package database

import (
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
