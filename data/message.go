package data

import "github.com/google/uuid"

type Message struct {
	Title string    `json:"title" binding:"required"`
	Body  string    `json:"body" binding:"required"`
	Cid   uuid.UUID `json:"cid" binding:"required"`
}
