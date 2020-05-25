package dao

import (
	"database/sql"
	"fmt"
)

var db *sql.DB

// SetDB ...
func SetDB(d *sql.DB) {
	db = d
}

// GetRevisionID ...
func GetRevisionID(hash []byte) (id int64) {

	result, err := db.Exec("INSERT IGNORE INTO revision SET hash = ?", hash)
	if err != nil {
		fmt.Println(`GetRevisionID fail:`, err)
		return
	}

	id, err = result.LastInsertId()
	if err == nil {
		return
	}

	return
}
