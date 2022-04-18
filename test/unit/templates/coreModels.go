package templates

import "github.com/qst-project/backend.git/pkg/core"

func GetCoreQuestion() *core.Question {
	return &core.Question{
		Statement: "Как дела?",
		Type:      0,
		Options: []string{
			"Хорошо",
			"Нормально",
			"Плохо",
		},
	}
}

func GetCoreQuestionnaire() *core.Questionnaire {
	return &core.Questionnaire{
		Ref:   "/testRef",
		Title: "Test Request",
		Questions: []*core.Question{
			{
				Id:        1,
				Statement: "Как дела?",
				Type:      0,
				Options: []string{
					"Хорошо",
					"Нормально",
					"Плохо",
				}},
		},
	}
}
