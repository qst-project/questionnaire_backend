package api

import (
	"github.com/qst-project/backend.git/pkg/usecase"
)

type QuestionnaireDelegate struct {
	usecase.QuestionnaireUseCase
}

func (s *QuestionnaireDelegate) CreateQuestionnaire(Questionnaire *Questionnaire) (ref string, err error) {
	if qst, err := Questionnaire.ToCore(); err == nil {
		ref, err = s.QuestionnaireUseCase.Invoke(qst)
	}
	return
}

func (s *QuestionnaireDelegate) GetQuestionnaire(ref string) (Questionnaire *Questionnaire, err error) {
	//qst, err := s.QuestionnaireUseCase.GetQuestionnaire(ref)
	return
}
