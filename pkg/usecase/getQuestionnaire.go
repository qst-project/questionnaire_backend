package usecase

import (
	"github.com/qst-project/backend.git/pkg/core"
	"github.com/qst-project/backend.git/pkg/gateway"
)

type GetQuestionnaireUseCaseImpl struct {
	gateway.QuestionnaireGateway
}

func (s *GetQuestionnaireUseCaseImpl) Invoke(ref string) (core.Questionnaire, error) {
	return s.QuestionnaireGateway.GetQuestionnaire(ref)
}
