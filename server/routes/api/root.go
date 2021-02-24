package api

import (
	"time"
)

// Post is a struct for dealing with new Post
type Post struct {
	ID    int64     `json:"id"`
	Title string    `json:"title"`
	Body  string    `json:"body"`
	Date  time.Time `json:"date"`
	IP    string    `json:"-"`
	Tags  []string  `json:"tags"`
}
