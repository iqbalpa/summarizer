package model

import "time"

type Summary struct {
	ID        string    `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Summary   string    `json:"summary"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// a summary belongs to one job
	JobId string
}

type SummaryRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}
