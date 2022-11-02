package database

import (
	"encoding/json"

	"github.com/dgraph-io/badger/v3"
)

func GetItem(db *badger.DB, key string, toSave interface{}) error {
	err := db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))

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

		err = json.Unmarshal(valCopy, toSave)

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func SaveItem(db *badger.DB, key string, data []byte) error {
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

func DeleteItem(db *badger.DB, key string) error {
	err := db.Update(func(txn *badger.Txn) error {
		err := txn.Delete([]byte(key))
		return err
	})

	if err != nil {
		return err
	}

	return err
}
