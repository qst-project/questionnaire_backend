package delegate

import (
	"github.com/qst-project/backend.git/pkg/usecase"
	"go.uber.org/fx"
)

type Module struct {
	fx.Out

	QuestionnaireDelegate
	HttpQuestionnaireDelegate
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

func HttpSetup(createQuestionnaireUseCase usecase.CreateQuestionnaireUseCase,
	getQuestionnaireUseCase usecase.GetQuestionnaireUseCase,
	updateQuestionnaireUseCase usecase.UpdateQuestionnaireUseCase) Module {
	return Module{
		HttpQuestionnaireDelegate: HttpQuestionnaireDelegate{
			createQuestionnaireUseCase,
			getQuestionnaireUseCase,
			updateQuestionnaireUseCase,
		},
	}
}
