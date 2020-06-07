package db

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"soulogic/pb"

	"github.com/dgraph-io/badger/v2"
	"google.golang.org/protobuf/proto"
)

var errItemNoChange = errors.New(`item no change`)

// ItemSet ...
func ItemSet(id uint32, rev *pb.Revision) (rid uint32, err error) {

	var revHash revHash
	revHash, err = setDBRevision(rev)
	if err != nil {
		return
	}

	itemMutex.Lock()
	defer itemMutex.Unlock()

	var item *pb.Item
	if id == 0 {
		item = itemNew(rev, revHash)
	} else {
		item, err = itemUpdate(id, rev, revHash)
		if err != nil {
			return
		}
	}

	err = setDBItem(item)
	j(`setDBItem`, err)

	rid = item.ID

	return
}

func itemNew(rev *pb.Revision, revHash revHash) (item *pb.Item) {

	item = &pb.Item{
		ID:           itemNextID,
		TsCreate:     ts(),
		RevisionHash: revHash[:],
		Revision:     rev,
	}

	itemNextID++
	itemPool[item.ID] = item

	return
}

func itemUpdate(id uint32, rev *pb.Revision, revHash revHash) (item *pb.Item, err error) {

	org, ok := itemPool[id]
	if !ok {
		err = fmt.Errorf(`item %d not found`, id)
		return
	}

	if bytes.Equal(org.RevisionHash, revHash[:]) {
		err = errItemNoChange
		return
	}

	item = &pb.Item{
		ID:           id,
		TsCreate:     org.TsCreate,
		TsRevise:     ts(),
		TsHide:       org.TsHide,
		RevisionHash: revHash[:],
		Revision:     rev,
	}

	return
}

func setDBItem(d *pb.Item) (err error) {

	if d.ID < 1 {
		err = errors.New(`empty item id`)
		return
	}

	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b, d.ID)

	k := append([]byte{byte(pb.Prefix_Item)}, b...)

	var tmpItem *pb.Item
	tmpItem = proto.Clone(d).(*pb.Item)
	tmpItem.Revision = nil

	var v []byte
	v, err = proto.Marshal(d)
	if err != nil {
		return
	}

	j(`save new item`, tmpItem)

	txn := db.NewTransaction(true)
	txn.Set(k, v)
	err = txn.Commit()

	return
}

func getDBItem(key []byte) (r *pb.Item, err error) {

	ab := []byte{byte(pb.Prefix_Item)}
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

	r = &pb.Item{}

	err = proto.Unmarshal(ab, r)
	if err != nil {
		r = nil
		return
	}

	// id := binary.LittleEndian.Uint32(key)

	return
}
