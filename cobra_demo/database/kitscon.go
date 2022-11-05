package database

import (
	"errors"
	"fmt"

	"github.com/dgraph-io/badger/v3"
	"github.com/google/uuid"
)

type KitsconList struct {
	Kitscons []uuid.UUID `json:"kitscons"`
}

type Kitscon struct {
	Id              uuid.UUID   `json:"id"`
	Name            string      `json:"name"`
	Desc            string      `json:"desc"`
	PresentationIds []uuid.UUID `json:"presentationIds"`
}

func GetKitscons(db *badger.DB) []Kitscon {
	// Get list of KitsCons
	var kitscons KitsconList
	err := GetItem(db, "kitscons", &kitscons)
	if errors.Is(err, badger.ErrKeyNotFound) {
		kitscons.Kitscons = []uuid.UUID{}
	}

	kitsconList := make([]Kitscon, len(kitscons.Kitscons))
	// For each uuid in list
	for i := range kitscons.Kitscons {
		// Get actual KitsCon from database
		var tempKitscon Kitscon
		err := GetItem(db, kitscons.Kitscons[i].String(), &tempKitscon)
		if err != nil {
			fmt.Printf("Error getting kitscon %s: %v\n", kitscons.Kitscons[i].String(), err)
		}
		kitsconList[i] = tempKitscon
	}

	return kitsconList
}
