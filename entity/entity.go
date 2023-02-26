package entity

import (
	"time"
)

type Entity struct {
	ID       uint       `json:"id,omitempty"`
	UpdateAt *time.Time `json:"update_at,omitempty"`
	CreateAt *time.Time `json:"create_at,omitempty"`
}
