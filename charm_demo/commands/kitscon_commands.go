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
	Id   uuid.UUID `json:"id"`
	Name string    `json:"name"`
	Desc string    `json:"desc"`
}

func (k Kitscon) Title() string       { return k.Name }
func (k Kitscon) Description() string { return k.Desc }
func (k Kitscon) FilterValue() string { return k.Name }

type KitsconsMsg []Kitscon

type KitsconAddedMsg bool

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
			fmt.Printf("Error when saving %v: %v", newKitscon, err)
			return KitsconAddedMsg(false)
		}

		return KitsconAddedMsg(true)
	}
}
