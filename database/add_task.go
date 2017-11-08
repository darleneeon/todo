package database

import (
	"encoding/json"
	"errors"

	"github.com/boltdb/bolt"
	"github.com/darleneeon/todo/model"
)

// AddTask saves task in the database.
func AddTask(t *model.Task) error {
	db, err := ConnectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("tasks"))

		// Generate ID for task.
		id, _ := b.NextSequence()
		t.ID = int(id)

		// Marshal task data into bytes
		buf, err := json.Marshal(t)
		if err != nil {
			err = errors.New("Marshalling task: " + err.Error())
			return err
		}

		// Persist bytes to tasks bucket
		return b.Put(itob(t.ID), buf)
	})
}
