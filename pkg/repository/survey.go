package repository

import (
	// "fmt"
	"github.com/skinnykaen/quesionnaire_backend.git/pkg/models"
	"github.com/skinnykaen/quesionnaire_backend.git/pkg/api"
)

type SurveyRepository struct {
	postgresClient *PostgresClient
}


func NewSurveyRepository(postgresClient *PostgresClient) *SurveyRepository {
	return &SurveyRepository{postgresClient: postgresClient}
}

func (r *SurveyRepository) GetSurvey(id string) (api.Survey, error) {
	survey := models.Questionnaire{}
	survey1 := api.Survey {
		Id: id,
		Title: "Test",
		Questions: []*api.Survey_Question{
			{Id: "1", Text: "Как дела?", Order: "1", Kind: "1"},
			{Id: "2", Text: "Какой сегодня день?", Order: "2", Kind: "3"},
		},
	}
	r.postgresClient.db.First(&survey, "1")
	// fmt.Println(survey)
	return survey1, nil
}
