package database

import (
	"time"

	"github.com/zachlatta/boolr/model"
)

const booleanGetByIDStmt = `SELECT id, created, updated, user_id, label, bool,
switch_count FROM booleans WHERE id = $1`

const booleanGetAllForUser = `SELECT id, created, updated, user_id, label,
bool, switch_count FROM booleans WHERE user_id = $1`

const booleanCreateStmt = `INSERT INTO booleans (created, updated, user_id,
label, bool, switch_count) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`

const booleanUpdateStmt = `UPDATE booleans SET updated=$1, label=$2, bool=$3, switch_count=$4 WHERE id=$5`

func GetBoolean(id int64) (*model.Boolean, error) {
	b := model.Boolean{}
	row := db.QueryRow(booleanGetByIDStmt, id)
	if err := row.Scan(&b.ID, &b.Created, &b.Updated, &b.UserID, &b.Label,
		&b.Bool, &b.SwitchCount); err != nil {
		return nil, err
	}
	return &b, nil
}

func GetBooleansForUser(userID int64) ([]*model.Boolean, error) {
	booleans := []*model.Boolean{}
	rows, err := db.Query(booleanGetAllForUser, userID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		b := model.Boolean{}
		if err := rows.Scan(&b.ID, &b.Created, &b.Updated, &b.UserID, &b.Label,
			&b.Bool, &b.SwitchCount); err != nil {
			return nil, err
		}

		booleans = append(booleans, &b)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return booleans, nil
}

func SaveBoolean(b *model.Boolean) error {
	saved := true
	if b.ID == 0 {
		b.Created = time.Now()
		saved = false
	}
	b.Updated = time.Now()

	if !saved {
		row := db.QueryRow(booleanCreateStmt, b.Created, b.Updated, b.UserID,
			b.Label, b.Bool, b.SwitchCount)
		if err := row.Scan(&b.ID); err != nil {
			return err
		}
	} else {
		if _, err := db.Exec(booleanUpdateStmt, b.Updated, b.Label, b.Bool,
			b.SwitchCount, b.ID); err != nil {
			return err
		}
	}
	return nil
}
