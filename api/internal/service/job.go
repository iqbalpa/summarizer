package service

import (
	"summarizer/internal/model"
	"summarizer/internal/repo"
)

type IJobService interface {
	GetJob(id string) (model.Job, error)
	CreateJob(title, content string) (model.Job, error)
}

type JobService struct {
	jr repo.IJobRepository
	ss ISummaryService
}

func NewJobService(jr repo.IJobRepository, ss ISummaryService) *JobService {
	return &JobService{
		jr: jr,
		ss: ss,
	}
}

func (js *JobService) GetJob(id string) (model.Job, error) {
	j, err := js.jr.GetJob(id)
	if err != nil {
		return model.Job{}, err
	}
	return j, nil
}

func (js *JobService) CreateJob(title, content string) (model.Job, error) {
	j, err := js.jr.CreateJob(model.Job{
		UserId: "219ea4f7-b72f-4f85-9aa4-65c335cbc985",
		Status: model.Pending,
	})
	if err != nil {
		return model.Job{}, err
	}
	_, err = js.ss.CreateSummary(title, content, j.ID)
	if err != nil {
		return model.Job{}, err
	}
	return j, nil
}
