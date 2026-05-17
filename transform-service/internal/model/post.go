package model

import "time"

type Post struct {
	ID          int       `bson:"_id"`
	Title       string    `bson:"title"`
	Body        string    `bson:"body"`
	Tags        []string  `bson:"tags"`
	Views       int       `bson:"views"`
	ProcessedAt time.Time `bson:"processed_at"`
}
