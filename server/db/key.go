package db

import (
	"encoding/binary"
	"soulogic/pb"
)

func key(a pb.Prefix_Enum, k interface{}) (ab []byte) {

	ab = []byte{byte(a)}

	switch a {

	case pb.Prefix_Item:
		b := make([]byte, 4)
		binary.LittleEndian.PutUint32(b, k.(uint32))
		ab = append(ab, b...)

	case pb.Prefix_Revision:
		ab = append(ab, k.([]byte)...)
	}

	return
}

// KeyItem ...
func KeyItem(id uint32) []byte {
	return key(pb.Prefix_Item, id)
}

// KeyRevision ...
func KeyRevision(hash []byte) []byte {
	return key(pb.Prefix_Revision, hash)
}
