package models

import (
	"time"
)

type Base struct {
	Id        int64     `json:"id" sql:",pk"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
