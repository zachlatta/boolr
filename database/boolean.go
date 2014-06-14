package database

import (
	"time"

	"github.com/zachlatta/boolr/model"
)

const booleanGetByIDStmt = `SELECT id, created, updated, user_id, label, bool,
switch_count FROM booleans WHERE id = $1`

const booleanCreateStmt = `INSERT INTO booleans (created, updated, user_id,
label, bool, switch_count) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`

func GetBoolean(id int64) (*model.Boolean, error) {
	b := model.Boolean{}
	row := db.QueryRow(booleanGetByIDStmt, id)
	if err := row.Scan(&b.ID, &b.Created, &b.Updated, &b.UserID, &b.Label,
		&b.Bool, &b.SwitchCount); err != nil {
		return nil, err
	}
	return &b, nil
}

func SaveBoolean(b *model.Boolean) error {
	if b.ID == 0 {
		b.Created = time.Now()
	}
	b.Updated = time.Now()

	row := db.QueryRow(booleanCreateStmt, b.Created, b.Updated, b.UserID,
		b.Label, b.Bool, b.SwitchCount)
	if err := row.Scan(&b.ID); err != nil {
		return err
	}
	return nil
}
