package repository

import (
	"github.com/qst-project/backend.git/pkg/models"
)

type Survey interface {
	SetSurvey(survey models.UtilitySurvey) (bool, error)
	GetSurvey(ref string) (models.UtilitySurvey, error)
}

type Repository struct {
	Survey
}

func NewRepository(postgresClient PostgresClient) *Repository {
	return &Repository{
		Survey: NewSurveyRepository(&postgresClient),
	}
}
