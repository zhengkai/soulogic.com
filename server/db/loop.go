package db

import (
	"encoding/hex"

	"github.com/dgraph-io/badger/v2"
)

// AllKey ...
func AllKey() {

	opt := badger.DefaultIteratorOptions
	opt.PrefetchSize = 10

	/*
		opt := &badger.IteratorOptions{
			PrefetchValue:  false,
			PrefetchSize:   10,
			AllVersion:     false,
			Prefix:         []byte(`k`),
			InternalAccess: false,
		}
	*/

	count := 0

	err := db.View(func(txn *badger.Txn) error {
		it := txn.NewIterator(opt)
		defer it.Close()
		for it.Rewind(); it.Valid(); it.Next() {

			key := it.Item().Key()

			j(hex.EncodeToString(key))

			GetRevision(key[1:])

			count++
		}
		return nil
	})

	j(`all key`, count, err)
}
