package database

import (
	"encoding/binary"
	"errors"
	"log"

	"github.com/boltdb/bolt"
)

func init() {
	db, err := ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create required bucket
	if err = db.Update(func(tx *bolt.Tx) error {
		_, err = tx.CreateBucketIfNotExists([]byte("tasks"))
		return err
	}); err != nil {
		log.Fatalf("Error while creating 'tasks' bucket: %s\n", err)
	}
}

// ConnectDB returns database connection.
func ConnectDB() (*bolt.DB, error) {
	db, err := bolt.Open("todo.db", 0600, nil)
	if err != nil {
		err = errors.New("Connecting to the database: " + err.Error())
		return nil, err
	}

	return db, nil
}

// itob returns an 8-byte big endian representation of v.
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}
