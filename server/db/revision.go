package db

import (
	"crypto/sha1"
	"soulogic/pb"

	"github.com/dgraph-io/badger/v2"
	"google.golang.org/protobuf/proto"
)

// SetRevision ...
func SetRevision(d *pb.Revision) {

	v, _ := proto.Marshal(d)

	hash := sha1.Sum(v)

	k := KeyRevision(hash[:])

	j(`SaveRevision`, len(hash), len(v))

	txn := db.NewTransaction(true)
	txn.Set(k, v)
	txn.Commit()

	item := &pb.Item{
		RevisionHash: hash[:],
	}

	j(item)
}

// GetRevision ...
func GetRevision(key []byte) (r *pb.Revision, err error) {

	ab := []byte{byte(pb.Prefix_Revision)}
	ab = append(ab, key...)

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
		return
	}

	j(r)

	return
}
