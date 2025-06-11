package model

import "time"

type Bookmark struct {
	ID        int       `json:"id"`
	URL       string    `json:"url"`
	Note      string    `json:"note"`
	Tags string `json:"tags"` // e.g., "go,backend,learning"
	CreatedAt time.Time `json:"created_at"`
}
