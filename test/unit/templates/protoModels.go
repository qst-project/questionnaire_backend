package templates

import (
	"github.com/qst-project/backend.git/pkg/api"
)

func GetProtoQuestion() *api.Question {
	return &api.Question{
		Statement: "Как дела?",
		Type:      api.Types_RADIO,
		Options: []string{
			"Хорошо",
			"Нормально",
			"Плохо",
		},
	}
}

func GetProtoQuestions() []*api.Question {
	return []*api.Question{{
		Statement: "Как дела?",
		Type:      api.Types_RADIO,
		Options: []string{
			"Хорошо",
			"Нормально",
			"Плохо",
		},
	},
		{
			Statement: "Cколько тебе лет?",
			Type:      api.Types_RADIO,
			Options: []string{
				"<20",
				">20",
			},
		},
	}
}

func GetProtoQuestionnaire() *api.Questionnaire {
	return &api.Questionnaire{
		Ref:       "/testRef",
		Title:     "Test Request",
		Questions: GetProtoQuestions(),
	}
}
