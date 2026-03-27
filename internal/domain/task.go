package domain

import "time"

type Task struct {
	ID          string
	ColumnID    string
	Header      string
	Description string
	Assignee    string
	Position    int64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
