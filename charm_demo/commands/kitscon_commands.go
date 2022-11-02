package commands

import (
	"encoding/json"
	"errors"
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/dgraph-io/badger/v3"
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
		k, _ := getListOfKitscons(db)
		kitsconList := make([]Kitscon, len(k.Kitscons))
		// For each uuid in list
		for i := range k.Kitscons {
			// Get actual KitsCon from database
			kitscon, err := getKitscon(db, k.Kitscons[i])
			if err != nil {
				fmt.Printf("Error getting kitscon %s: %v\n", k.Kitscons[i].String(), err)
			}
			kitsconList[i] = kitscon
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
		err = saveItem(db, newKitscon.Id.String(), marshalled)
		if err != nil {
			fmt.Printf("Error when saving %v: %v", newKitscon, err)
			return KitsconAddedMsg(false)
		}

		// Fetch list of KitsCons, add new kitscon to it
		kitscons, _ := getListOfKitscons(db)
		kitscons.Kitscons = append(kitscons.Kitscons, newKitscon.Id)
		marshalled, err = json.Marshal(kitscons)
		if err != nil {
			fmt.Printf("Error when marshalling %v: %v", kitscons, err)
			return KitsconAddedMsg(false)
		}

		// Update list in database
		err = saveItem(db, "kitscons", marshalled)
		if err != nil {
			fmt.Printf("Error when saving %v: %v", newKitscon, err)
			return KitsconAddedMsg(false)
		}

		return KitsconAddedMsg(true)
	}
}

// ------------- HELPERS ----------------
func getListOfKitscons(db *badger.DB) (KitsconList, error) {
	var kitscons KitsconList
	err := db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte("kitscons"))

		if errors.Is(err, badger.ErrKeyNotFound) {
			kitscons = KitsconList{Kitscons: []uuid.UUID{}}
			return nil
		}

		var valCopy []byte
		err = item.Value(func(val []byte) error {
			valCopy = append([]byte{}, val...)

			return nil
		})

		if err != nil {
			return err
		}

		err = json.Unmarshal(valCopy, &kitscons)

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return kitscons, err
	}

	return kitscons, nil
}

func getKitscon(db *badger.DB, kitsconId uuid.UUID) (Kitscon, error) {
	var kitscon Kitscon
	err := db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(kitsconId.String()))
		if err != nil {
			return err
		}

		var valCopy []byte
		err = item.Value(func(val []byte) error {
			valCopy = append([]byte{}, val...)

			return nil
		})

		if err != nil {
			return err
		}

		err = json.Unmarshal(valCopy, &kitscon)

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return kitscon, err
	}

	return kitscon, nil
}

func saveItem(db *badger.DB, key string, data []byte) error {
	err := db.Update(func(txn *badger.Txn) error {
		e := badger.NewEntry([]byte(key), data)
		err := txn.SetEntry(e)
		return err
	})

	if err != nil {
		return err
	}

	return nil
}
