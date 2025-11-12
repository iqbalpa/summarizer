package repo

import (
	"fmt"
	"summarizer/internal/model"

	"gorm.io/gorm"
)

type ISummaryRepository interface {
	GetSummary(id string) (model.Summary, error)
	GetAllSummaries() ([]model.Summary, error)
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
	fmt.Printf("Before create: %+v\n", s)
	result := sr.db.Create(&s)
	if result.Error != nil {
		return model.Summary{}, fmt.Errorf("failed to create a new summary")
	}
	fmt.Printf("After create: %+v\n", s)
	return s, nil
}

func (sr *SummaryRepository) GetAllSummaries() ([]model.Summary, error) {
	var summaries []model.Summary
	result := sr.db.Find(&summaries)
	if result.Error != nil {
		return []model.Summary{}, fmt.Errorf("failed to retrieve all summaries")
	}
	return summaries, nil
}
