package commands

import (
	"encoding/json"
	"errors"
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/dgraph-io/badger/v3"
	"github.com/doverstav/kitscon222/charm_demo/database"
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

func (k Kitscon) Title() string       { return k.Name }
func (k Kitscon) Description() string { return k.Desc }
func (k Kitscon) FilterValue() string { return k.Name }

type KitsconsMsg []Kitscon

type KitsconSelectedMsg Kitscon

type KitsconAddedMsg bool

type KitsconRemovedMsg bool

func GetKitscons(db *badger.DB) tea.Cmd {
	return func() tea.Msg {
		// Get list of KitsCons
		var kitscons KitsconList
		err := database.GetItem(db, "kitscons", &kitscons)
		if errors.Is(err, badger.ErrKeyNotFound) {
			kitscons.Kitscons = []uuid.UUID{}
		}

		kitsconList := make([]Kitscon, len(kitscons.Kitscons))
		// For each uuid in list
		for i := range kitscons.Kitscons {
			// Get actual KitsCon from database
			var tempKitscon Kitscon
			err := database.GetItem(db, kitscons.Kitscons[i].String(), &tempKitscon)
			if err != nil {
				fmt.Printf("Error getting kitscon %s: %v\n", kitscons.Kitscons[i].String(), err)
			}
			kitsconList[i] = tempKitscon
		}

		return KitsconsMsg(kitsconList)
	}
}

func SaveKitscon(db *badger.DB, name string, description string) tea.Cmd {
	return func() tea.Msg {
		// Create new KitsCon struct
		newKitscon := Kitscon{Id: uuid.New(), Name: name, Desc: description}
		marshalled, err := json.Marshal(newKitscon)
		if err != nil {
			fmt.Printf("Error when marshalling %v: %v", newKitscon, err)
			return KitsconAddedMsg(false)
		}

		// Save to database
		err = database.SaveItem(db, newKitscon.Id.String(), marshalled)
		if err != nil {
			fmt.Printf("Error when saving %v: %v", newKitscon, err)
			return KitsconAddedMsg(false)
		}

		// Fetch list of KitsCons, add new kitscon to it
		var kitscons KitsconList
		err = database.GetItem(db, "kitscons", &kitscons)
		if errors.Is(err, badger.ErrKeyNotFound) {
			kitscons.Kitscons = []uuid.UUID{}
		}
		kitscons.Kitscons = append(kitscons.Kitscons, newKitscon.Id)

		marshalled, err = json.Marshal(kitscons)
		if err != nil {
			fmt.Printf("Error when marshalling %v: %v", kitscons, err)
			return KitsconAddedMsg(false)
		}

		// Update list in database
		err = database.SaveItem(db, "kitscons", marshalled)
		if err != nil {
			fmt.Printf("Error when saving %v: %v", kitscons, err)
			return KitsconAddedMsg(false)
		}

		return KitsconAddedMsg(true)
	}
}

func DeleteKitscon(db *badger.DB, kitsconToRemove uuid.UUID, presentations []uuid.UUID) tea.Cmd {
	return func() tea.Msg {
		// Remove presentation under kitscon
		for _, presentationId := range presentations {
			err := database.DeleteItem(db, presentationId.String())
			if err != nil {
				fmt.Printf("Failed to delete presentation %s: %v", presentationId.String(), err)
			}
		}

		// Remove actual kitscon
		err := database.DeleteItem(db, kitsconToRemove.String())
		if err != nil {
			fmt.Printf("Could not delete kitscon %s: %v", kitsconToRemove.String(), err)
			return KitsconRemovedMsg(false)
		}

		// Get list of KitsCons
		var kitscons KitsconList
		err = database.GetItem(db, "kitscons", &kitscons)
		if errors.Is(err, badger.ErrKeyNotFound) {
			kitscons.Kitscons = []uuid.UUID{}
		}

		// Remove deleted kitscon from it
		filteredKitscons := []uuid.UUID{}
		for _, kitsconId := range kitscons.Kitscons {
			if kitsconId != kitsconToRemove {
				filteredKitscons = append(filteredKitscons, kitsconId)
			}
		}
		kitscons.Kitscons = filteredKitscons

		marshalled, err := json.Marshal(kitscons)
		if err != nil {
			fmt.Printf("Error when marshalling %v: %v", kitscons, err)
			return KitsconRemovedMsg(false)
		}

		// Update list in database
		err = database.SaveItem(db, "kitscons", marshalled)
		if err != nil {
			fmt.Printf("Error when saving %v: %v", kitscons, err)
			return KitsconRemovedMsg(false)
		}

		return KitsconRemovedMsg(true)
	}
}

func KitsconSelected(kitscon Kitscon) tea.Cmd {
	return func() tea.Msg {
		//fmt.Printf("Selected %v", kitscon)
		return KitsconSelectedMsg(kitscon)
	}
}
