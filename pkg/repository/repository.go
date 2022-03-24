package repository

import 	"github.com/skinnykaen/quesionnaire_backend.git/pkg/api"

type Survey interface {
	SetSurvey(survey api.Survey) (bool, string,error)
	GetSurvey(ref string) (api.Survey, error)
}

type Repository struct {
	Survey
}

func NewRepository(postgresClient PostgresClient) *Repository {
	return &Repository{
		Survey: NewSurveyRepository(&postgresClient),
	}
}