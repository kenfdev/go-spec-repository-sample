package domain

import "time"

type Invoice struct {
	ID      string
	Total   int
	DueDate time.Time
}
