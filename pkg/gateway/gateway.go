package gateway

import (
	"github.com/qst-project/backend.git/pkg/core"
)

type QuestionnaireGateway interface {
	GetQuestionnaire(ref string) (core.Questionnaire, error)
	CreateQuestionnaire(Questionnaire core.Questionnaire) (bool, error)
	DeleteQuestionnaire(ref string) (bool, error)
	UpdateQuestionnaire(Questionnaire core.Questionnaire) (bool, error)
}

type Module struct {
	QuestionnaireGateway
}

func NewGatewayModule(postgresClient PostgresClient) *Module {
	return &Module{
		QuestionnaireGateway: NewQuestionnaireGateway(&postgresClient),
	}
}
