package database

import (
	"encoding/json"
	"log"

	"github.com/boltdb/bolt"
	"github.com/darleneeon/todo/model"
)

// AddTask saves task in the database.
func AddTask(t *model.Task) error {
	db, err := ConnectDB()
	if err != nil {
		log.Fatalf("Error while connecting to the database: %s\n", err)
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
			return err
		}

		// Persist bytes to tasks bucket
		return b.Put(itob(t.ID), buf)
	})
}
