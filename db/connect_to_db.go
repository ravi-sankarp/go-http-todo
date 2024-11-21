package dbState

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type dbState struct {
	Error     bool
	Connected bool
}

var Db *sql.DB

func ConnectToDb() {
	connstring := "user=postgres dbname=postgres password=123456 host=host.docker.internal port=5434 sslmode=disable"
	db, err := sql.Open("postgres", connstring)
	if err != nil {
		fmt.Println(err)
		DbState.Error = true
		panic(err)
	}
	Db = db
	DbState.Connected = true
}

var DbState = &dbState{
	Error:     false,
	Connected: false,
}
