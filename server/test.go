package main

import (
	"fmt"
	"soulogic/db"
	"soulogic/pb"
)

func test() {

	ab := db.KeyItem(uint32(256))

	j(`test`, len(ab), ab)

	v := &pb.Revision{
		Format: pb.PostFormat_Markdown,
		Raw:    fmt.Sprintf(`test rpg %d`, ts()),
	}

	null(v)
	// db.SetRevision(v)
}
