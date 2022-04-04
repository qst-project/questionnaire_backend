package gateway

import (
	"github.com/qst-project/backend.git/pkg/core"
)

type QuestionnaireGateway interface {
	GetQuestionnaire(ref string) (core.Questionnaire, error)
	CreateQuestionnaire(Questionnaire core.Questionnaire) (string, error)
	DeleteQuestionnaire(ref string) (bool, error)
	UpdateQuestionnaire(Questionnaire core.Questionnaire) (bool, error)
}

type Module struct {
	QuestionnaireGateway
}

func Setup(postgresClient *PostgresClient) *Module {
	return &Module{
		QuestionnaireGateway: &QuestionnaireGatewayImpl{postgresClient: postgresClient},
	}
}
