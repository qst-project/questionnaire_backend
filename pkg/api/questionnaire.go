package api

import (
	"github.com/qst-project/backend.git/pkg/usecase"
)

type QuestionnaireDelegate struct {
	usecase.CreateQuestionnaireUseCase
	usecase.GetQuestionnaireUseCase
	usecase.UpdateQuestionnaireUseCase
}

func (s *QuestionnaireDelegate) CreateQuestionnaire(Questionnaire *Questionnaire) (ref string, err error) {
	qst := Questionnaire.ToCore()
	ref, err = s.CreateQuestionnaireUseCase.Invoke(qst)
	return
}

func (s *QuestionnaireDelegate) GetQuestionnaire(ref string) (Questionnaire, error) {
	var protoQuestionnaire Questionnaire
	qst, err := s.GetQuestionnaireUseCase.Invoke(ref)
	if err != nil {
		return protoQuestionnaire, err
	}
	protoQuestionnaire.FromCore(&qst)
	return protoQuestionnaire, nil
}

func (s *QuestionnaireDelegate) UpdateQuestionnaire(Questionnaire *Questionnaire) (err error) {
	qst := Questionnaire.ToCore()
	err = s.UpdateQuestionnaireUseCase.Invoke(qst)
	return
}
