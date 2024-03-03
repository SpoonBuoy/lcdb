package db

import (
	"sync"

	"github.com/spoonbuoy/lcdb/config"
)

type Database struct {
	Config config.DBConfig
	Mu     sync.Mutex
}

func (db *Database) Connect() (Store, error) {
	return db, nil
}

func (db *Database) Close() error {
	return nil
}
func (db *Database) GetClient() error {
	return nil
}
