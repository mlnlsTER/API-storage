package memdb

import "training/30_DB_Tasks/pkg/storage/postgres"

type DB []postgres.Task

func (db DB) Tasks(int, int) ([]postgres.Task, error) {
	return db, nil
}

func (db DB) NewTask(postgres.Task) (int, error) {
	return 0, nil
}
