package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"soulogic/dao"
)

func initDB() {

	d, err := sql.Open("mysql", "soulogic:password@/soulogic")
	if err != nil {
		j(`mysql fail`, err)
		return
	}

	dao.SetDB(d)

	j(`GetRevisionID`, dao.GetRevisionID([]byte{0, 0, 1}))
}
