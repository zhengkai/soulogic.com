package db

import (
	"soulogic/pb"
	"time"

	"github.com/dgraph-io/badger/v2"
)

// AllKey ...
func allKey() {

	var opt badger.IteratorOptions

	count := 0

	t := time.Now()

	opt = badger.IteratorOptions{
		PrefetchValues: false,
		PrefetchSize:   100,
		Reverse:        false,
		AllVersions:    false,
		Prefix:         []byte{2},
	}

	err := db.View(func(txn *badger.Txn) (err error) {
		it := txn.NewIterator(opt)
		defer it.Close()
		for it.Rewind(); it.Valid(); it.Next() {

			key := it.Item().Key()

			err = initSwitchKey(key)
			if err != nil {
				j(`err`, err)
			}
			count++
		}
		return
	})

	opt = badger.IteratorOptions{
		PrefetchValues: false,
		PrefetchSize:   100,
		Reverse:        false,
		AllVersions:    false,
		Prefix:         []byte{1},
	}

	err = db.View(func(txn *badger.Txn) (err error) {
		it := txn.NewIterator(opt)
		defer it.Close()
		for it.Rewind(); it.Valid(); it.Next() {

			key := it.Item().Key()

			err = initSwitchKey(key)
			if err != nil {
				j(`err`, err)
			}
			count++
		}
		return
	})

	d := time.Since(t)

	itemNextID++

	j(`all key`, count, err, d)
	j(`itemNextID`, itemNextID)
	j(`item pool`, len(itemPool))
}

func initSwitchKey(key []byte) (err error) {
	t := pb.Prefix_Enum(key[0])

	// j(`initSwitchKey`, hex.EncodeToString(key))

	subKey := key[1:]

	switch t {

	case pb.Prefix_Revision:

		hash := getRevHash(subKey)

		var r *pb.Revision
		r, err = getDBRevision(hash)
		if err != nil {
			j(`error when load revision`, err.Error())
			return
		}

		revPool[hash] = r

	case pb.Prefix_Item:

		var r *pb.Item
		r, err = getDBItem(subKey)
		if err != nil {
			j(`error when load item`, err.Error())
			return
		}

		if itemNextID < r.ID {
			itemNextID = r.ID
		}

		itemPool[r.ID] = r

	default:
		j(`unknown key`, string(key))
	}

	return
}
