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

func (s *SurveyService) GetSurvey(ref string) (api.Survey, error) {
	// transform to proto model
	return  s.repo.GetSurvey(ref)
}

func (s *SurveyService) SetSurvey(survey api.Survey) (bool, string, error) {
	// transform to gorm model
	var utilitySurvey UtilitySurvey
	utilitySurvey.From(survey)
	return  s.repo.SetSurvey(utilitySurvey)
}

func (em *UtilitySurvey) From(gRRCModel api.Survey) {
	var questionnaire models.Questionnaire
	var question models.Question
	var radioPossibleAnswer models.RadioPossibleAnswer
	var radioPossibleAnswer models.RadioPossibleAnswer
	var textPossibleAnswer models.TextPossibleAnswer


	questionnaire.From(survey)

	em.Questionnaire.Id = quesionnaire.ID
	em.Questionnaire.Title = quesionnaire.Title
	em.Questionnaire.Ref = quesionnaire.Ref

	for index, question := range survey.Questions{
		question.From(survey.Questions[index])
		em.Questions.append(em.Questions, question)
	}

	// -----------------------------------------------------------------------------

	for index, radioPossibleAnswer := range survey.RadioPossibleAnswers{
		radioPossibleAnswer.From(survey.RadioPossibleAnswers[index])
		em.RadioPossibleAnswers.append(em.RadioPossibleAnswers, radioPossibleAnswer)
	}

	for index, checkboxPossibleAnswer := range survey.CheckboxPossibleAnswers{
		checkboxPossibleAnswer.From(survey.CheckboxPossibleAnswers[index])
		em.CheckboxPossibleAnswers.append(em.CheckboxPossibleAnswers, checkboxPossibleAnswer)
	}

	for index, textPossibleAnswer := range survey.TextPossibleAnswers{
		textPossibleAnswer.From(survey.TextPossibleAnswers[index])
		em.TextPossibleAnswers.append(em.TextPossibleAnswers, textPossibleAnswer)
	}

	// ----------------------------------------------------------------------------

	for index, radioAnswer := range survey.RadioAnswers{
		radioAnswer.From(survey.RadioAnswers[index])
		em.RadioAnswers.append(em.RadioAnswers, radioAnswer)
	}
	for index, checkboxAnswer := range survey.CheckboxAnswers{
		checkboxAnswer.From(survey.CheckboxAnswers[index])
		em.CheckboxAnswers.append(em.CheckboxAnswers, checkboxAnswer)
	}
	for index, textAnswer := range survey.TextAnswers{
		textAnswer.From(survey.TextAnswers[index])
		em.TextAnswers.append(em.TextAnswers, textAnswer)
	}
}

// вспомогательный struct 
type UtilitySurvey struct {
	Questionnaire *models.Questionnaire
	Questions []*models.Question
	RadioPossibleAnswers []*models.RadioPossibleAnswer
	CheckboxPossibleAnswers []*models.CheckboxPossibleAnswer
	TextPossibleAnswers []*models.TextPossibleAnswer
	RadioAnswers []*models.RadioAnswer
	CheckboxAnswers []*models.CheckboxAnswer
	TextAnswers []*TextAnswer
} 