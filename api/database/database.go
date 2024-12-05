package database

import (
	"bytes"
	"encoding/gob"

	"github.com/dgraph-io/badger/v4"
)

var DB *badger.DB

func Exists(key string) bool {
	var exists bool
	err := DB.View(func(txn *badger.Txn) error {
		_, err := txn.Get([]byte(key))
		exists = err == nil
		return nil
	})

	if err != nil {
		return false
	}

	return exists
}

func Get(key string) ([]byte, error) {
	var data []byte
	err := DB.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))
		if err != nil {
			return err
		}

		err = item.Value(func(val []byte) error {
			data = append([]byte{}, val...)
			return nil
		})

		return err
	})

	return data, err
}

func UpdateOrInsert(key string, value interface{}) error {
	err := DB.Update(func(txn *badger.Txn) error {
		var data bytes.Buffer
		enc := gob.NewEncoder(&data)

		err := enc.Encode(value)
		if err != nil {
			return err
		}

		err = txn.Set([]byte(key), data.Bytes())
		return err
	})

	return err
}
