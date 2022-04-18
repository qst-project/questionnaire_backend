package api

import (
	"github.com/qst-project/backend.git/pkg/usecase"
	"go.uber.org/fx"
)

type Module struct {
	fx.Out

	QuestionnaireDelegate
}

func Setup(createQuestionnaireUseCase usecase.CreateQuestionnaireUseCase,
	getQuestionnaireUseCase usecase.GetQuestionnaireUseCase,
	updateQuestionnaireUseCase usecase.UpdateQuestionnaireUseCase) Module {
	return Module{
		QuestionnaireDelegate: QuestionnaireDelegate{
			createQuestionnaireUseCase,
			getQuestionnaireUseCase,
			updateQuestionnaireUseCase,
		},
	}
}
