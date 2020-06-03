package db

import (
	"soulogic/pb"
	"sync"
)

var (
	itemNextID uint32

	itemPool  = make(map[uint32]*pb.Item)
	itemMutex sync.Mutex

	revPool  = make(map[revHash]*pb.Revision)
	revMutex sync.Mutex
)

type revHash [20]byte
