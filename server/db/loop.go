package db

import (
	"encoding/hex"
	"soulogic/pb"

	"github.com/dgraph-io/badger/v2"
)

// AllKey ...
func allKey() {

	opt := badger.DefaultIteratorOptions
	opt.PrefetchSize = 1000

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

	err := db.View(func(txn *badger.Txn) (err error) {
		it := txn.NewIterator(opt)
		defer it.Close()
		for it.Rewind(); it.Valid(); it.Next() {

			key := it.Item().Key()

			j(hex.EncodeToString(key))

			var hash revHash
			copy(hash[:], key[1:])

			var r *pb.Revision
			r, err = getDBRevision(hash)
			if err != nil {
				return
			}

			revPool[hash] = r

			count++
		}
		return
	})

	j(`all key`, count, err)
}
