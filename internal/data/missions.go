package data

import (
	"time"
)

type Mission struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Who       string    `json:"who"`
	What      string    `json:"what"`
	Where     string    `json:"where"`
	When      string    `json:"when"`
	Why       string    `json:"why"`
}
