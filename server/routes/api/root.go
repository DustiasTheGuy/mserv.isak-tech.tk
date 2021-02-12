package api

import (
	"time"
)

// Post is a struct for dealing with new Post
type Post struct {
	ID   uint      `json:"_id"`
	Body string    `json:"body"`
	Date time.Time `json:"date"`
	IP   string    `json:"ip"`
}
