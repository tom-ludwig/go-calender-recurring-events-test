package data

import (
	"github.com/google/uuid"
)

type Events struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	StartTime   string    `json:"start_time"`
	EndTime     string    `json:"end_time"`
	RRule       string    `json:"rrule"`
	ExRule      string    `json:"exrule"`
}
