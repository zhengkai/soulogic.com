package db

func dbSet(k, v []byte) error {
	txn := db.NewTransaction(true)
	txn.Set(k, v)
	return txn.Commit()
}
