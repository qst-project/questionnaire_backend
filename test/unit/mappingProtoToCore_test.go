package unit

import (
	"github.com/qst-project/backend.git/pkg/api"
	"github.com/qst-project/backend.git/pkg/core"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMappingProtoQuestionnaireToCore(t *testing.T) {
	protoQuestionnaire := &api.Questionnaire{
		Ref:   "/testRef",
		Title: "Test Request",
		Questions: []*api.Question{
			{Id: "1", Statement: "Как дела?", Type: 0, Options: []string{
				"Хорошо",
				"Нормально",
				"Плохо",
			}},
		},
	}
	coreQuestionnaire := protoQuestionnaire.ToCore()

	assert.Equal(t, protoQuestionnaire.Ref, coreQuestionnaire.Ref)
	assert.Equal(t, protoQuestionnaire.Title, coreQuestionnaire.Title)
	assert.Equal(t, protoQuestionnaire.Questions[0].Statement, coreQuestionnaire.Questions[0].Statement)
}

func TestMappingProtoQuestionnaireFromCore(t *testing.T) {
	coreQuestionnaire := &core.Questionnaire{
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
	var protoQuestionnaire api.Questionnaire
	protoQuestionnaire.FromCore(coreQuestionnaire)

	assert.Equal(t, protoQuestionnaire.Ref, coreQuestionnaire.Ref)
	assert.Equal(t, protoQuestionnaire.Title, coreQuestionnaire.Title)
	assert.Equal(t, protoQuestionnaire.Questions[0].Statement, coreQuestionnaire.Questions[0].Statement)
}

func TestMappingProtoQuestionToCore(t *testing.T) {
	protoQuestion := &api.Question{
		Statement: "Как дела?",
		Type:      0,
		Options: []string{
			"Хорошо",
			"Нормально",
			"Плохо",
		},
	}

	coreQuestion := protoQuestion.ToCore()

	assert.Equal(t, protoQuestion.Statement, coreQuestion.Statement)
	assert.Equal(t, uint(protoQuestion.Type), uint(coreQuestion.Type))
	assert.Equal(t, protoQuestion.Options[0], coreQuestion.Options[0])
}

func TestMappingProtoQuestionFromCore(t *testing.T) {
	coreQuestion := &core.Question{
		Statement: "Как дела?",
		Type:      1,
		Options: []string{
			"Хорошо",
			"Нормально",
			"Плохо",
		},
	}
	var protoQuestion api.Question
	protoQuestion.FromCore(coreQuestion)

	assert.Equal(t, protoQuestion.Statement, coreQuestion.Statement)
	assert.Equal(t, uint(protoQuestion.Type), uint(coreQuestion.Type))
	assert.Equal(t, protoQuestion.Options[0], coreQuestion.Options[0])
}
