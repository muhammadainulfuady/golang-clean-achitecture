package helper

import (
	"database/sql"
)

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func CommitOrRollback(tx *sql.Tx) {
	err := recover()
	if err != nil {
		tx.Rollback()
		panic(err)
	} else {
		tx.Commit()
	}
}
