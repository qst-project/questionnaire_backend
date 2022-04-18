package templates

import (
	"github.com/qst-project/backend.git/pkg/proto"
)

func GetProtoQuestion() *proto.Question {
	return &proto.Question{
		Statement: "Как дела?",
		Type:      proto.Types_RADIO,
		Options: []string{
			"Хорошо",
			"Нормально",
			"Плохо",
		},
	}
}

func GetProtoQuestions() []*proto.Question {
	return []*proto.Question{{
		Statement: "Как дела?",
		Type:      proto.Types_RADIO,
		Options: []string{
			"Хорошо",
			"Нормально",
			"Плохо",
		},
	},
		{
			Statement: "Cколько тебе лет?",
			Type:      proto.Types_RADIO,
			Options: []string{
				"<20",
				">20",
			},
		},
	}
}

func GetProtoQuestionnaire() *proto.Questionnaire {
	return &proto.Questionnaire{
		Ref:       "testRef",
		Title:     "Test Request",
		Questions: GetProtoQuestions(),
	}
}
