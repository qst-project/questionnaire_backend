package templates

import (
	"github.com/qst-project/backend.git/pkg/gateway"
)

func GetDbQuestion() *gateway.QuestionDB {
	return &gateway.QuestionDB{
		Statement: "Как дела?",
		Type:      0,
		Order:     1,
	}
}

func GetDbQuestions() []*gateway.QuestionDB {
	return []*gateway.QuestionDB{
		{
			Type:      1,
			Statement: "Как дела?",
			Order:     1,
		},
		{
			Type:      1,
			Statement: "Сколько тебе лет?",
			Order:     2,
		},
	}
}

func GetDbOptions() []*gateway.OptionsDB {
	return []*gateway.OptionsDB{
		{
			QuestionId: 1,
			Text:       "Хорошо",
		},
		{
			QuestionId: 1,
			Text:       "Нормально",
		},
		{
			QuestionId: 1,
			Text:       "Плохо",
		},
	}
}

func GetDbQuestionnaire() *gateway.QuestionnaireDB {
	return &gateway.QuestionnaireDB{
		Ref:   "testRef",
		Title: "Test Request",
	}
}
