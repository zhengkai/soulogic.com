package main

import (
	"math/rand"
	"os"
	"path/filepath"
	"soulogic/db"
	"soulogic/imp"
	"time"
)

var pwd string

func main() {
	jw(`start`)

	initMain()

	defer closeLog()
	initLog()

	err := imp.Start()
	if err != nil {
		j(`mysql fail:`, err)
	}

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
