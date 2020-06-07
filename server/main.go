package main

import (
	"math/rand"
	"os"
	"path/filepath"
	"soulogic/db"
	"time"
)

var pwd string

func main() {
	jw(`start`)

	initMain()

	defer closeLog()
	initLog()

	initConfig()

	db.Start(pwd + `/badger`)

	server()

	test()
}

func initMain() (err error) {

	pwd, err = filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return
	}

	rand.Seed(time.Now().UTC().UnixNano() - 1585133837641794656)

	return
}
