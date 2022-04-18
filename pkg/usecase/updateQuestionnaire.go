package usecase

import (
	"github.com/qst-project/backend.git/pkg/core"
	"github.com/qst-project/backend.git/pkg/gateway"
)

type UpdateQuestionnaireUseCaseImpl struct {
	gateway.QuestionnaireGateway
}

func (s *UpdateQuestionnaireUseCaseImpl) Invoke(Questionnaire core.Questionnaire) error {
	return s.QuestionnaireGateway.UpdateQuestionnaire(Questionnaire)
}
