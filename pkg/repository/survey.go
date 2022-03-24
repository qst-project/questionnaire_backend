package repository

import (
	// "fmt"
	"github.com/skinnykaen/quesionnaire_backend.git/pkg/models"
	"github.com/skinnykaen/quesionnaire_backend.git/pkg/api"
	"github.com/skinnykaen/quesionnaire_backend.git/pkg/service"
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

func (r *SurveyRepository) SetSurvey(survey service.UtilitySurvey) (bool, string, error) {
	r.postgresClient.db.Create(survey.Quesionnaire)
	if r.postgresClient.db.Create(survey.Quesionnaire).Error != nil {
		//log.Fatal(err)
		return false, nil, r.postgresClient.db.Create(survey.Quesionnaire).Error 
	}
	r.postgresClient.db.Create(survey.Questions)
	r.postgresClient.db.Create(survey.RadioPossibleAnswers)
	r.postgresClient.db.Create(survey.CheckboxPossibleAnswers)
	r.postgresClient.db.Create(survey.TextPossibleAnswers)
	r.postgresClient.db.Create(survey.RadioAnswers)
	r.postgresClient.db.Create(survey.CheckboxAnswers)
	r.postgresClient.db.Create(survey.TextAnswers)
	return true, "ref", nil
}
