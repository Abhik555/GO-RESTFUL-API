package models

import (
	"time"

	"github.com/Abhik555/GO-RESTFUL-API/db"
)

type Event struct {
	ID          int
	Name        string `binding:"required"`
	Description string `binding:"required"`
	Location    string `binding:"required"`
	CreatedAt   time.Time
	UserID      string
}

func (event *Event) Save() error {
	query := `
	INSERT INTO EVENTS(NAME , DESCRIPTION , LOCATION , CREATEDAT , USERID)
	VALUES(?,?,?,?,?)
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	res, err := stmt.Exec(event.Name, event.Description, event.Location, time.Now(), event.UserID)
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	event.ID = int(id)
	return err
}

func GetAllEvents() ([]Event, error) {
	query := `
	SELECT * FROM EVENTS
	`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.CreatedAt, &event.UserID)

		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}

func GetEventbyID(id int64) (*Event, error) {
	query := "SELECT * FROM EVENTS WHERE ID=?"

	row := db.DB.QueryRow(query, id)

	var event Event

	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.CreatedAt, &event.UserID)

	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (e Event) Update() error {
	query := `
	UPDATE EVENTS
	SET NAME = ? , DESCRIPTION = ? , LOCATION = ? , CREATEDAT = ? , USERID = ?
	WHERE ID = ?
	`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.Name, e.Description, e.Location, e.CreatedAt, e.UserID, e.ID)
	return err
}

func (e Event) Delete() error {
	query := "DELETE FROM EVENTS WHERE ID = ?"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.ID)

	return err
}

func (e Event) Register(userID string) error {
	query := "INSERT INTO REGISTRATIONS(EVENTID , USERID) VALUES(? , ?)"

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.ID, userID)
	return err
}

func (e Event) CancelRegistration(userID string) error {
	query := "DELETE FROM REGISTRATIONS WHERE EVENTID = ? AND USERID = ?"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.ID, userID)
	return err
}
