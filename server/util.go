package main

import "time"

func ts() uint32 {
	return uint32(time.Now().Unix())
}

func tsm() uint64 {
	return uint64(time.Now().UnixNano() / 1000000)
}

func null(_ interface{}) {
}
