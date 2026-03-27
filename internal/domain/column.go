package domain

import "time"

type Column struct {
	ID        string
	Name      string
	BoardID   string
	Position  int64
	CreatedAt time.Time
	UpdatedAt time.Time
}
