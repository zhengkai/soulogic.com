package db

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"sort"
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
	if err != nil {
		return
	}
	// j(`setDBItem`, err)

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
	v, err = proto.Marshal(tmpItem)
	if err != nil {
		return
	}

	txn := db.NewTransaction(true)
	txn.Set(k, v)
	err = txn.Commit()

	if err != nil {
		return
	}

	itemPool[d.ID] = d

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

	revHash := getRevHash(r.RevisionHash)
	r.Revision, _ = revPool[revHash]

	j(`rev`, revHash, r.Revision)

	// id := binary.LittleEndian.Uint32(key)

	return
}

// ItemRecent ...
func ItemRecent() (list []*pb.Item) {

	var idList []uint32

	itemMutex.Lock()
	for _, v := range itemPool {
		idList = append(idList, v.ID)
	}

	sort.Slice(idList, func(i, j int) bool {
		return idList[i] > idList[j]
	})

	if len(idList) > 1000 {
		idList = idList[0:1000]
	}

	for _, id := range idList {
		list = append(list, itemPool[id])
	}
	itemMutex.Unlock()

	return
}
