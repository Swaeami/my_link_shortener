package domain

import "time"

type URL struct {
	ID        int       `json:"id" binding:"required"`
	Original  string    `json:"original" binding:"required"`
	Short     string    `json:"short" binding:"required"`
	CreatedAt time.Time `json:"created_at" binding:"required"`
}
