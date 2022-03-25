package controller

import (
	"database/sql"
	"log"
)

func connect() *sql.DB {
	db, err := sql.Open("mysql", "root:1234@tcp(localhost:3306)/explr_martini?parseTime=true&loc=Asia%2FJakarta")

	if err != nil {
		log.Fatal(err)
	}
	return db
}
