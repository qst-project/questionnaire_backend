package delegate

import (
	"github.com/qst-project/backend.git/pkg/httpModels"
	"github.com/qst-project/backend.git/pkg/proto"
	"github.com/qst-project/backend.git/pkg/usecase"
)

type HttpQuestionnaireDelegate struct {
	usecase.CreateQuestionnaireUseCase
	usecase.GetQuestionnaireUseCase
	usecase.UpdateQuestionnaireUseCase
}

func (s *HttpQuestionnaireDelegate) CreateQuestionnaire(Questionnaire *httpModels.Questionnaire) (ref string, err error) {
	qst := Questionnaire.ToCore()
	ref, err = s.CreateQuestionnaireUseCase.Invoke(qst)
	return
}

func (s *HttpQuestionnaireDelegate) GetQuestionnaire(ref string) (httpModels.Questionnaire, error) {
	var httpQuestionnaire httpModels.Questionnaire
	qst, err := s.GetQuestionnaireUseCase.Invoke(ref)
	httpQuestionnaire.FromCore(&qst)
	if err != nil {
		return httpQuestionnaire, err
	}
	return httpQuestionnaire, nil
}

func (s *HttpQuestionnaireDelegate) UpdateQuestionnaire(Questionnaire *proto.Questionnaire) (err error) {
	qst := Questionnaire.ToCore()
	err = s.UpdateQuestionnaireUseCase.Invoke(qst)
	return
}
