package api

import (
	"github.com/qst-project/backend.git/pkg/usecase"
)

type QuestionnaireDelegate struct {
	usecase.CreateQuestionnaireUseCase
}

func (s *QuestionnaireDelegate) CreateQuestionnaire(Questionnaire *Questionnaire) (ref string, err error) {
	if qst, err := Questionnaire.ToCore(); err == nil {
		ref, err = s.CreateQuestionnaireUseCase.Invoke(qst)
	}
	return
}
