package main

import (
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

var pwd string

func main() {

	initMain()

	initDB()
}

func initMain() (err error) {

	pwd, err = filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return
	}

	rand.Seed(time.Now().UTC().UnixNano() - 1585133837641794656)

	return
}
