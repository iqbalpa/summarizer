package repo

import (
	"fmt"
	"summarizer/internal/model"

	"gorm.io/gorm"
)

type IJobRepository interface {
	GetJob(id string) (model.Job, error)
	CreateJob(model.Job) (model.Job, error)
}

type JobRepository struct {
	db *gorm.DB
}

func NewJobRepository(db *gorm.DB) *JobRepository {
	return &JobRepository{
		db: db,
	}
}

func (jr *JobRepository) GetJob(id string) (model.Job, error) {
	var job model.Job
	result := jr.db.Where("id = ?", id).First(&job)
	if result.Error != nil {
		return model.Job{}, fmt.Errorf("job with id %s not found", id)
	}
	return job, nil
}

func (jr *JobRepository) CreateJob(j model.Job) (model.Job, error) {
	result := jr.db.Create(&j)
	if result.Error != nil {
		return model.Job{}, fmt.Errorf("failed to create a new job")
	}
	return j, nil
}
