package helper

import "database/sql"

func CommitRollBack(tx *sql.Tx) {
	err := recover()
	if err != nil {
		err := tx.Rollback()
		ErrorHandle(err)
		panic(err)
	} else {
		err := tx.Commit()
		ErrorHandle(err)

	}
}
