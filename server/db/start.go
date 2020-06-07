package db

import "github.com/dgraph-io/badger/v2"

var db *badger.DB

var itemIDSerial uint32

// Start ...
func Start(path string) (err error) {

	db, err = badger.Open(badger.DefaultOptions(path))
	if err != nil {
		return
	}

	allKey()

	table := db.Tables(false)
	j(`table`, table)

	return
}

// End ...
func End() {
	db.Close()
}
