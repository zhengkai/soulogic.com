package main

import "soulogic/db"

func initDB() (err error) {

	path := pwd + `/badger`

	err = db.Start(path)
	if err != nil {
		return
	}

	db.AllKey()

	/*
		db, err = badger.Open(badger.DefaultOptions())
		if err != nil {
			return
		}
		defer db.Close()

		table := db.Tables(false)
		for _, v := range table {
			j(`table`, v)
		}

		key := []byte(`key`)

		seq, err := db.GetSequence(key, 10)
		if err != nil {
			return
		}

		defer seq.Release()
		i := 10
		for {
			i--
			if i <= 0 {
				break
			}
			num, err := seq.Next()

			j(`next`, num, err)
			if err != nil {
				break
			}
		}
	*/

	/*
		d, err := sql.Open("mysql", "soulogic:password@/soulogic")
		if err != nil {
			j(`mysql fail`, err)
			return
		}

		dao.SetDB(d)

		j(`GetRevisionID`, dao.GetRevisionID([]byte{0, 0, 1}))
	*/

	return
}
