package repository

import 	"github.com/skinnykaen/quesionnaire_backend.git/pkg/api"

type Survey interface {
	// SetSurvey(id uint) (error)
	GetSurvey(id string) (api.Survey, error)
}

type Repository struct {
	Survey
}

func NewRepository(postgresClient PostgresClient) *Repository {
	return &Repository{
		Survey: NewSurveyRepository(&postgresClient),
	}
}