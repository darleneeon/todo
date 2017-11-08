package database

import (
	"log"

	"github.com/boltdb/bolt"
)

// DeleteTask deletes task from the database.
func DeleteTask(id int) error {
	db, err := ConnectDB()
	if err != nil {
		log.Fatalf("Error while connecting to the database: %s\n", err)
	}
	defer db.Close()

	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("tasks"))

		return b.Delete(itob(id))
	})
}
