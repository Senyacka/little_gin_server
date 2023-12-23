package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)


func init() {
	db, err := sql.Open("mysql", "root:Password1234!@/MySQL80")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Verify that the connection is made
	err = db.Ping()
	if err != nil {
		panic(err)
	}
}
