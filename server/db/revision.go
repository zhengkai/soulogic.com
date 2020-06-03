package db

import (
	"crypto/sha1"
	"soulogic/pb"

	"github.com/dgraph-io/badger/v2"
	"google.golang.org/protobuf/proto"
)

// SetRevision ...
func setDBRevision(d *pb.Revision) (hash revHash, err error) {

	v, _ := proto.Marshal(d)

	hash = sha1.Sum(v)

	revMutex.Lock()

	_, ok := revPool[hash]
	if ok {
		revMutex.Unlock()
		return
	}

	defer revMutex.Unlock()

	k := append([]byte{byte(pb.Prefix_Revision)}, hash[:]...)

	txn := db.NewTransaction(true)
	txn.Set(k, v)
	err = txn.Commit()
	if err != nil {
		return
	}

	return
}

func getDBRevision(key revHash) (r *pb.Revision, err error) {

	ab := []byte{byte(pb.Prefix_Revision)}
	ab = append(ab, key[:]...)

	var item *badger.Item
	txn := db.NewTransaction(false)
	item, err = txn.Get(ab)
	if err != nil {
		return
	}
	txn.Commit()

	ab, err = item.ValueCopy(nil)
	if err != nil {
		return
	}

	r = &pb.Revision{}

	err = proto.Unmarshal(ab, r)
	if err != nil {
		r = nil
		return
	}

	return
}
