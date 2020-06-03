package db

import (
	"fmt"
	"time"
)

var j = fmt.Println

func ts() uint32 {
	return uint32(time.Now().Unix())
}

func null(_ interface{}) {
}
