package usecase

import (
	"github.com/qst-project/backend.git/pkg/core"
	"github.com/qst-project/backend.git/pkg/gateway"
)

type QuestionnaireUseCase interface {
	GetQuestionnaire(ref string) (core.Questionnaire, error)
	CreateQuestionnaire(Questionnaire *core.Questionnaire) (bool, string, error)
	DeleteQuestionnaire(ref string) (bool, error)
	UpdateQuestionnaire(Questionnaire *core.Questionnaire) (bool, error)
}

type Module struct {
	QuestionnaireUseCase
}

func NewUseCaseModule(repos *gateway.Module) *Module {
	return &Module{
		QuestionnaireUseCase: NewQi(repos),
	}
}
