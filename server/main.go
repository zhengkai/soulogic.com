package main

import (
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"github.com/dgraph-io/badger/v2"
)

var pwd string

func main() {

	initMain()

	db, err := badger.Open(badger.DefaultOptions(pwd + "/badger"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}

func initMain() (err error) {

	pwd, err = filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return
	}

	rand.Seed(time.Now().UTC().UnixNano() - 1585133837641794656)

	return
}
