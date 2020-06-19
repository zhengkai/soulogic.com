package imp

import (
	"database/sql"
	"fmt"
	"log"
	"soulogic/pb"

	_ "github.com/go-sql-driver/mysql"
)

var j = fmt.Println
var db *sql.DB

// Start ...
func Start() (err error) {

	db, err = sql.Open("mysql", "root:U9kMZodPS7jT3Ty@/soulogic_legacy")
	if err != nil {
		return
	}
	defer db.Close()

	rows, err := db.Query(`SELECT item_id, revision_id, ts_create FROM item`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id uint32
		var revisionID uint32
		var ts uint32

		err = rows.Scan(&id, &revisionID, &ts)
		if err != nil {
			return
		}

		j(id, revisionID, ts)

		item := &pb.Item{
			Revision: &pb.Revision{
				Format: pb.PostFormat_Markdown,
				Raw:    ``,
			},
		}

		db.ItemImport(item)

		err = findRow(revisionID)
		if err != nil {
			return
		}
	}

	return
}

func findRow(id uint32) (err error) {

	j(`findRow`, id)

	row := db.QueryRow(`SELECT content FROM revision WHERE revision_id = ?`, id)
	if row == nil {
		err = fmt.Errorf(`select fail %d`, id)
		return
	}

	var ab []byte

	err = row.Scan(&ab)
	if err != nil {
		return
	}

	j(len(ab))

	return
}
