package delegate

import (
	"github.com/qst-project/backend.git/pkg/api"
	"github.com/qst-project/backend.git/pkg/usecase"
)

type QuestionnaireDelegate struct {
	usecase.CreateQuestionnaireUseCase
}

func (s *QuestionnaireDelegate) CreateQuestionnaire(Questionnaire *api.Questionnaire) (ref string, err error) {
	if qst, err := Questionnaire.ToCore(); err == nil {
		ref, err = s.CreateQuestionnaireUseCase.Invoke(qst)
	}
	return
}
