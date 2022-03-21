package service

import (
	"github.com/skinnykaen/quesionnaire_backend.git/pkg/repository"
	"github.com/skinnykaen/quesionnaire_backend.git/pkg/api"
)

type SurveyService struct {
	repo repository.Survey
}

func NewSurveyService(repo *repository.Repository) *SurveyService {
	return &SurveyService{repo: repo}
}

func (s *SurveyService) GetSurvey(id string) (api.Survey, error) {
	return  s.repo.GetSurvey(id)
}