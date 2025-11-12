package model

import "time"

type JobStatus string

const (
	Pending    JobStatus = "pending"
	InProgress JobStatus = "in-progress"
	Completed  JobStatus = "completed"
	Archived   JobStatus = "archived"
	Deleted    JobStatus = "deleted"
)

type Job struct {
	ID        string    `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Status    JobStatus `json:"job_status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// a job belongs to one user
	UserId string

	// a job has one summary
	Summary Summary
}

type JobRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}
