package badger

import "github.com/dgraph-io/badger/v4"

func NewBadgerDB() (*badger.DB, error) {

	db, err := badger.Open(badger.DefaultOptions("").WithInMemory(true))
	if err != nil {
		return nil, err
	}

	return db, err
}
