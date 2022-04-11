package usecase

import (
	"github.com/qst-project/backend.git/pkg/core"
	"github.com/qst-project/backend.git/pkg/gateway"
	"go.uber.org/fx"
)

type CreateQuestionnaireUseCase interface {
	Invoke(Questionnaire core.Questionnaire) (string, error)
}

type Module struct {
	fx.Out

	CreateQuestionnaireUseCase
}

func Setup(questionnaireGateway gateway.QuestionnaireGateway) Module {
	return Module{
		CreateQuestionnaireUseCase: &CreateQuestionnaireUseCaseImpl{questionnaireGateway},
	}
}
