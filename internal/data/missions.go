package data

import (
	"time"
)

type Mission struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Who       string    `json:"who,omitempty"`
	What      string    `json:"what,omitempty"`
	Where     string    `json:"where,omitempty"`
	When      string    `json:"when,omitempty"`
	Why       string    `json:"why,omitempty"`
}
