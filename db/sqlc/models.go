// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"time"
)

type Url struct {
	ID        int64     `json:"id"`
	LongUrl   string    `json:"long_url"`
	ShortUrl  string    `json:"short_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
