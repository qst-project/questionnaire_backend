package usecase

import (
	"github.com/qst-project/backend.git/pkg/core"
	"github.com/qst-project/backend.git/pkg/gateway"
	"go.uber.org/fx"
)

type CreateQuestionnaireUseCase interface {
	Invoke(Questionnaire core.Questionnaire) (string, error)
}

type GetQuestionnaireUseCase interface {
	Invoke(ref string) (core.Questionnaire, error)
}

type UpdateQuestionnaireUseCase interface {
	Invoke(Questionnaire core.Questionnaire) error
}

type Module struct {
	fx.Out

	CreateQuestionnaireUseCase
	GetQuestionnaireUseCase
	UpdateQuestionnaireUseCase
}

func Setup(questionnaireGateway gateway.QuestionnaireGateway) Module {
	return Module{
		CreateQuestionnaireUseCase: &CreateQuestionnaireUseCaseImpl{questionnaireGateway},
		GetQuestionnaireUseCase:    &GetQuestionnaireUseCaseImpl{questionnaireGateway},
		UpdateQuestionnaireUseCase: &UpdateQuestionnaireUseCaseImpl{questionnaireGateway},
	}
}
