package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Computer struct {
	ID    int64  `json:"id"`
	Owner string `json:"owner"`
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
	// catch to error.
}

func addUser(db *sql.DB, desription string, owner string) {
	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("insert into computers (description,owner) values (?,?)")
	_, err := stmt.Exec(desription, owner)
	checkError(err)
	tx.Commit()
}

func main() {
	fmt.Println("Hello, It Manager")

	db, _ := sql.Open("sqlite3", "database/computers.db")
	db.Exec("create table if not exists computers (id integer not null primary key autoincreament,owner text,description text)")

	fmt.Println("Your database is ready!")

	addUser(db, "Super Beast", "Norbert")

	fmt.Println("Your database is ready!")
}
