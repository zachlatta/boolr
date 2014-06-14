package database

import (
	"database/sql"
	"time"

	"github.com/zachlatta/boolr/model"
)

const userGetByIDStmt = `SELECT id, created, updated, username, password FROM
users WHERE id = $1`

const userGetByUsernameStmt = `SELECT id, created, updated, username, password
FROM users WHERE username ilike $1`

const userCreateStmt = `INSERT INTO users (created, updated, username,
password) VALUES ($1, $2, $3, $4) RETURNING id`

func GetUser(id int64) (*model.User, error) {
	u := new(model.User)
	row := db.QueryRow(userGetByIDStmt, id)
	if err := row.Scan(&u.ID, &u.Created, &u.Updated, &u.Username,
		&u.Password); err != nil {
		return nil, err
	}
	return u, nil
}

func GetUserByUsername(username string) (*model.User, error) {
	u := new(model.User)
	row := db.QueryRow(userGetByUsernameStmt, username)
	if err := row.Scan(&u.ID, &u.Created, &u.Updated, &u.Username,
		&u.Password); err != nil {
		return nil, err
	}
	return u, nil
}

func SaveUser(u *model.User) error {
	if u.ID == 0 {
		_, err := GetUserByUsername(u.Username)
		if err == nil {
			return model.ErrInvalidUserUsername
		} else if err != sql.ErrNoRows && err != nil {
			return err
		}

		u.Created = time.Now()
	}
	u.Updated = time.Now()

	rows, err := db.Query(userCreateStmt, u.Created, u.Updated, u.Username,
		u.Password)
	if err != nil {
		return err
	}
	for rows.Next() {
		if err := rows.Scan(&u.ID); err != nil {
			return err
		}
	}
	if err := rows.Err(); err != nil {
		return err
	}

	return nil
}
