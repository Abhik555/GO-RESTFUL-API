package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("Could not connet to database.")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {

	createUsersTable := `
	CREATE TABLE IF NOT EXISTS USERS(
	ID INTEGER PRIMARY KEY AUTOINCREMENT,
	EMAIL TEXT NOT NULL UNIQUE,
	PASSWORD TEXT NOT NULL
	)
	`
	_, err := DB.Exec(createUsersTable)

	if err != nil {
		panic("Could not create USER table.")

	}

	createEventsTable := `
	CREATE TABLE IF NOT EXISTS EVENTS(
	ID INTEGER PRIMARY KEY AUTOINCREMENT,
	NAME TEXT NOT NULL,
	DESCRIPTION TEXT NOT NULL,
	LOCATION TEXT NOT NULL,
	CREATEDAT DATETIME NOT NULL, 
	USERID TEXT,
	FOREIGN KEY(USERID) REFERENCES USERS(ID)
	)
	`

	_, err = DB.Exec(createEventsTable)

	if err != nil {
		fmt.Println(err)
		panic("Could not create EVENTS table.")
	}

	createRegisterationsTable := `
	CREATE TABLE IF NOT EXISTS REGISTRATIONS(
	ID INTEGER PRIMARY KEY AUTOINCREMENT,
	EVENTID INTEGER,
	USERID TEXT,
	FOREIGN KEY(EVENTID) REFERENCES EVENTS(ID),
	FOREIGN KEY(USERID) REFERENCES USERS(ID)
	)
	`

	_, err = DB.Exec(createRegisterationsTable)

	if err != nil {
		fmt.Println(err)
		panic("Could not create REG table.")
	}
}
