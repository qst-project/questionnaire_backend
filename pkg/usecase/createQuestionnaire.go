package usecase

import (
	"github.com/qst-project/backend.git/pkg/core"
	"github.com/qst-project/backend.git/pkg/gateway"
)

type CreateQuestionnaireUseCaseImpl struct {
	gateway.QuestionnaireGateway
}

func (s *CreateQuestionnaireUseCaseImpl) Invoke(Questionnaire core.Questionnaire) (string, error) {
	return s.QuestionnaireGateway.CreateQuestionnaire(Questionnaire)
}
