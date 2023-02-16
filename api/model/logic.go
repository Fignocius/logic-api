package model

import (
	"encoding/json"
	"time"

	uuid "github.com/satori/go.uuid"
)

type Logic struct {
	ID            uuid.UUID `json:"id" db:"id"`
	Expression    string    `json:"expression" db:"expression"`
	ExpresionCode string    `json:"-" db:"expression_code"`
	CreatedAt     time.Time `json:"-" db:"created_at"`
	UpdatedAt     time.Time `json:"-" db:"updated_at"`
}

func (m Logic) String() string {
	b, err := json.Marshal(m)
	if err != nil {
		return string("")
	}
	return string(b)
}
