package api

import (
	"github.com/qst-project/backend.git/pkg/usecase"
	"go.uber.org/fx"
)

type Module struct {
	fx.Out

	QuestionnaireDelegate
}

func Setup(QuestionnaireUseCase usecase.QuestionnaireUseCase) Module {
	return Module{
		QuestionnaireDelegate: QuestionnaireDelegate{QuestionnaireUseCase},
	}
}
