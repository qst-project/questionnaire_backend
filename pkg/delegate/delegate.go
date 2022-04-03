package delegate

import (
	"github.com/qst-project/backend.git/pkg/api"
	"github.com/qst-project/backend.git/pkg/gateway"
)

type Delegate interface {
	GetQuestionnaire(ref string) (*api.Questionnaire, error)
	CreateQuestionnaire(Questionnaire *api.Questionnaire) (string, error)
	DeleteQuestionnaire(ref string) (bool, error)
	UpdateQuestionnaire(Questionnaire *api.Questionnaire) (bool, error)
}

type Module struct {
	Delegate
}

func NewDelegateModule(repos *gateway.Module) *Module {
	return &Module{
		Delegate: NewDelegate(repos),
	}
}
