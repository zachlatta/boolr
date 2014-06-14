package model

import (
	"encoding/json"
	"errors"
	"io"
	"time"
)

var (
	ErrInvalidBooleanLabel = errors.New("invalid label")
)

type Boolean struct {
	ID          int64     `json:"id"`
	Created     time.Time `json:"created"`
	Updated     time.Time `json:"updated"`
	UserID      int64     `json:"user_id"`
	Label       string    `json:"label,omitempty"`
	Bool        bool      `json:"bool"`
	SwitchCount int64     `json:"switch_count"`
}

func NewBoolean(jsonReader io.Reader, userID int64) (*Boolean, error) {
	var b Boolean
	if err := json.NewDecoder(jsonReader).Decode(&b); err != nil &&
		err.Error() != "EOF" {
		return nil, err
	}

	if err := b.validate(); err != nil {
		return nil, err
	}

	boolean := Boolean{
		UserID:      userID,
		Label:       b.Label,
		Bool:        false,
		SwitchCount: 0,
	}

	return &boolean, nil
}

func (b *Boolean) Switch() {
	b.Bool = !b.Bool
	b.SwitchCount++
}

func (b *Boolean) validate() error {
	switch {
	case len(b.Label) > 255:
		return ErrInvalidBooleanLabel
	default:
		return nil
	}
}
