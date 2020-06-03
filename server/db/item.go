package db

import (
	"bytes"
	"errors"
	"fmt"
	"soulogic/pb"

	"google.golang.org/protobuf/proto"
)

var errItemNoChange = errors.New(`item no change`)

// ItemSet ...
func ItemSet(id uint32, rev *pb.Revision) (err error) {

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

	v, _ := proto.Marshal(item)
	k := append([]byte{byte(pb.Prefix_Revision)}, revHash[:]...)

	txn := db.NewTransaction(true)
	txn.Set(k, v)
	err = txn.Commit()
	if err != nil {
		return
	}

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
		err = fmt.Errorf(`item %d not found`, item)
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
