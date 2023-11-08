package badger

import "github.com/dgraph-io/badger/v4"

func NewBadgerDB() (*badger.DB, error) {
	opts := badger.Options{
		Dir:      "",
		Logger:   nil,
		InMemory: true,
	}
	db, err := badger.Open(opts)
	if err != nil {
		return nil, err
	}

	return db, err
}
