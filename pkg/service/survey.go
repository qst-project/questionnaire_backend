package service

import (
	"github.com/qst-project/backend.git/pkg/api"
	"github.com/qst-project/backend.git/pkg/models"
	"github.com/qst-project/backend.git/pkg/repository"
)

type SurveyService struct {
	repo repository.Survey
}

func NewSurveyService(repo *repository.Repository) *SurveyService {
	return &SurveyService{repo: repo}
}

func (s *SurveyService) GetSurvey(ref string) (*api.Survey, error) {
	utilitySurvey, err := s.repo.GetSurvey(ref)
	apiSurvey := utilitySurvey.GetgRPCModel()
	return &apiSurvey, err
}

func (s *SurveyService) SetSurvey(survey *api.Survey) (bool, string, error) {
	var utilitySurvey models.UtilitySurvey
	utilitySurvey.From(survey)
	result, err := s.repo.SetSurvey(utilitySurvey)
	return result, "generate ref", err
}

func (s *SurveyService) DeleteSurvey(ref string) (bool, error) {
	result, err := s.repo.DeleteSurvey(ref)
	if err != nil {
		return false, err
	}
	return result, nil
}
