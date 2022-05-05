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
		Id:        "1",
		Statement: "Как дела?",
		Type:      proto.Types_RADIO,
		Options: []string{
			"Хорошо",
			"Нормально",
			"Плохо",
		},
	},
		{
			Id:        "2",
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
		Id:        1,
		Ref:       "testRef",
		Title:     "Test Request",
		Questions: GetProtoQuestions(),
	}
}
