package database

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/dgraph-io/badger/v3"
	"github.com/google/uuid"
)

type ErrNoSuchKitscon struct{}

func (e *ErrNoSuchKitscon) Error() string {
	return "could not find kitscon"
}

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

func GetKitsconByName(db *badger.DB, kitsconName string) (Kitscon, error) {
	kitscons := GetKitscons(db)

	for _, kitscon := range kitscons {
		if kitscon.Name == kitsconName {
			return kitscon, nil
		}
	}

	return Kitscon{}, &ErrNoSuchKitscon{}
}

func SaveKitscon(db *badger.DB, name string, description string) error {
	// Create new KitsCon struct
	newKitscon := Kitscon{Id: uuid.New(), Name: name, Desc: description}
	marshalled, err := json.Marshal(newKitscon)
	if err != nil {
		fmt.Printf("Error when marshalling %v: %v", newKitscon, err)
		return err
	}

	// Save to database
	err = SaveItem(db, newKitscon.Id.String(), marshalled)
	if err != nil {
		fmt.Printf("Error when saving %v: %v", newKitscon, err)
		return err
	}

	// Fetch list of KitsCons, add new kitscon to it
	var kitscons KitsconList
	err = GetItem(db, "kitscons", &kitscons)
	if errors.Is(err, badger.ErrKeyNotFound) {
		kitscons.Kitscons = []uuid.UUID{}
	}
	kitscons.Kitscons = append(kitscons.Kitscons, newKitscon.Id)

	marshalled, err = json.Marshal(kitscons)
	if err != nil {
		fmt.Printf("Error when marshalling %v: %v", kitscons, err)
		return err
	}

	// Update list in database
	err = SaveItem(db, "kitscons", marshalled)
	if err != nil {
		fmt.Printf("Error when saving %v: %v", kitscons, err)
		return err
	}

	return nil

}
