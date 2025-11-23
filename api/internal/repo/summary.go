package repo

import (
	"fmt"
	"summarizer/internal/model"

	"gorm.io/gorm"
)

type ISummaryRepository interface {
	GetSummary(id string) (model.Summary, error)
	GetAllSummaries(userId string) ([]model.Summary, error)
	CreateSummary(summary model.Summary) (model.Summary, error)
}

type SummaryRepository struct {
	db *gorm.DB
}

func NewSummaryRepository(db *gorm.DB) *SummaryRepository {
	return &SummaryRepository{
		db: db,
	}
}

func (sr *SummaryRepository) GetSummary(id string) (model.Summary, error) {
	var s model.Summary
	result := sr.db.Where("id = ?", id).First(&s)
	if result.Error != nil {
		return model.Summary{}, fmt.Errorf("summary id %s not found", id)
	}
	return s, nil
}

func (sr *SummaryRepository) CreateSummary(s model.Summary) (model.Summary, error) {
	result := sr.db.Create(&s)
	if result.Error != nil {
		return model.Summary{}, fmt.Errorf("failed to create a new summary")
	}
	return s, nil
}

func (sr *SummaryRepository) GetAllSummaries(userId string) ([]model.Summary, error) {
	// Get all job ids belong to UserId
	var jobIds []string
	result := sr.db.Model(&model.Job{}).
		Select("id").
		Where("user_id = ?", userId).
		Find(&jobIds)
	if result.Error != nil {
		return []model.Summary{}, fmt.Errorf("failed to retrieve all job ids")
	}
	// get all summaries belong to userId, based on the jobIds
	var summaries []model.Summary
	result = sr.db.Where("job_id IN ?", jobIds).Find(&summaries)
	if result.Error != nil {
		return []model.Summary{}, fmt.Errorf("failed to retrieve all summaries")
	}
	return summaries, nil
}
