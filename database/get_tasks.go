package database

import (
	"encoding/json"
	"errors"

	"github.com/boltdb/bolt"
	"github.com/darleneeon/todo/model"
)

// GetTasks return slice of tasks.
func GetTasks() ([]model.Task, error) {
	var tasks []model.Task

	db, err := ConnectDB()
	if err != nil {
		return []model.Task{}, err
	}
	defer db.Close()

	if err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("tasks"))
		c := b.Cursor()

		cap := b.Sequence()
		tasks = make([]model.Task, 0, cap)

		var task model.Task
		for k, v := c.First(); k != nil; k, v = c.Next() {
			task, err = unmarshal(v)
			if err != nil {
				return err
			}

			tasks = append(tasks, task)
		}

		return nil
	}); err != nil {
		return []model.Task{}, err
	}

	return tasks, nil
}

func unmarshal(v []byte) (model.Task, error) {
	var task model.Task

	if err := json.Unmarshal(v, &task); err != nil {
		err = errors.New("Unmarshalling tasks: " + err.Error())
		return model.Task{}, err
	}

	return task, nil
}
