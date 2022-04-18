package delegate

import (
	"github.com/qst-project/backend.git/pkg/core"
	"github.com/qst-project/backend.git/pkg/proto"
	"github.com/qst-project/backend.git/pkg/usecase"
)

type QuestionnaireDelegate struct {
	usecase.CreateQuestionnaireUseCase
	usecase.GetQuestionnaireUseCase
	usecase.UpdateQuestionnaireUseCase
}

func (s *QuestionnaireDelegate) CreateQuestionnaire(Questionnaire *proto.Questionnaire) (ref string, err error) {
	qst := Questionnaire.ToCore()
	ref, err = s.CreateQuestionnaireUseCase.Invoke(qst)
	return
}

func (s *QuestionnaireDelegate) GetQuestionnaire(ref string) (core.Questionnaire, error) {
	qst, err := s.GetQuestionnaireUseCase.Invoke(ref)
	if err != nil {
		return qst, err
	}
	return qst, nil
}

func (s *QuestionnaireDelegate) UpdateQuestionnaire(Questionnaire *proto.Questionnaire) (err error) {
	qst := Questionnaire.ToCore()
	err = s.UpdateQuestionnaireUseCase.Invoke(qst)
	return
}
