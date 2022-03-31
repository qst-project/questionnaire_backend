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
	if err != nil {
		return nil, err
	}
	apiSurvey := utilitySurvey.GetgRPCModel()
	return &apiSurvey, nil
}

func (s *SurveyService) CreateSurvey(survey *api.Survey) (bool, string, error) {
	var utilitySurvey models.UtilitySurvey
	utilitySurvey.From(survey)
	result, err := s.repo.CreateSurvey(utilitySurvey)
	if err != nil {
		return false, "", err
	}
	return result, "generate ref", nil
}

func (s *SurveyService) DeleteSurvey(ref string) (bool, error) {
	result, err := s.repo.DeleteSurvey(ref)
	if err != nil {
		return false, err
	}
	return result, nil
}

func (s *SurveyService) UpdateSurvey(survey *api.Survey) (bool, error) {
	var utilitySurvey models.UtilitySurvey
	utilitySurvey.From(survey)
	result, err := s.repo.UpdateSurvey(utilitySurvey)
	return result, err
}
