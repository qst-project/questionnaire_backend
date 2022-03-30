package service

import (
	"github.com/qst-project/backend.git/pkg/api"
	"github.com/qst-project/backend.git/pkg/repository"
)

type Survey interface {
	GetSurvey(id string) (*api.Survey, error)
	SetSurvey(survey *api.Survey) (bool, string, error)
	// Update
	// Delete
}

type Service struct {
	Survey
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Survey: NewSurveyService(repos),
	}
}
