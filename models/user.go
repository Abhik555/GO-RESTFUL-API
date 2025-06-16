package models

import (
	"errors"

	"github.com/abhik555/EventsAPI/db"
	"github.com/abhik555/EventsAPI/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error {
	query := `
	INSERT INTO USERS(EMAIL , PASSWORD)
	VALUES(?,?)
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	pswd, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}

	res, err := stmt.Exec(u.Email, pswd)
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	u.ID = id
	return err
}

func (u *User) ValidateCredentials() error {
	query := "SELECT ID,PASSWORD FROM USERS WHERE EMAIL=?"
	row := db.DB.QueryRow(query, u.Email)

	var retriedpass string
	err := row.Scan(&u.ID, &retriedpass)

	if err != nil {
		return errors.New("Invalid credentials.")
	}

	isValid := utils.CheckpasswordHash(u.Password, retriedpass)

	if !isValid {
		return errors.New("Invalid credentials.")
	}

	return nil

}
