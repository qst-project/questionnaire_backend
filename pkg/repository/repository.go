package repository

import (
	"github.com/qst-project/backend.git/pkg/models"
)

type Survey interface {
	GetSurvey(ref string) (models.UtilitySurvey, error)
	CreateSurvey(survey models.UtilitySurvey) (bool, error)
	DeleteSurvey(ref string) (bool, error)
	UpdateSurvey(survey models.UtilitySurvey) (bool, error)
}

type Repository struct {
	Survey
}

func NewRepository(postgresClient PostgresClient) *Repository {
	return &Repository{
		Survey: NewSurveyRepository(&postgresClient),
	}
}
