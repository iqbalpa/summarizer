package service

import (
	"summarizer/internal/model"
	"summarizer/internal/repo"
)

type ISummaryService interface {
	GetSummary(id string) (model.Summary, error)
	CreateSummary(title, content, jobId string) (model.Summary, error)
	GetAllSummaries() ([]model.Summary, error)
}

type SummaryService struct {
	sr repo.ISummaryRepository
}

func NewSummaryService(sr repo.ISummaryRepository) *SummaryService {
	return &SummaryService{
		sr: sr,
	}
}

func (ss *SummaryService) GetSummary(id string) (model.Summary, error) {
	s, err := ss.sr.GetSummary(id)
	if err != nil {
		return model.Summary{}, err
	}
	return s, nil
}

func (ss *SummaryService) CreateSummary(title, content, jobId string) (model.Summary, error) {
	s, err := ss.sr.CreateSummary(model.Summary{
		Title:   title,
		Content: content,
		JobId:   jobId,
	})
	if err != nil {
		return model.Summary{}, err
	}
	return s, nil
}

func (ss *SummaryService) GetAllSummaries() ([]model.Summary, error) {
	s, err := ss.sr.GetAllSummaries()
	if err != nil {
		return []model.Summary{}, nil
	}
	return s, nil
}
