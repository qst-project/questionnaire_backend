package unit

import (
	"github.com/qst-project/backend.git/pkg/core"
	"github.com/qst-project/backend.git/pkg/gateway"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMappingDbQuestionnaireToCore(t *testing.T) {
	dbQuestionnaire := &gateway.QuestionnaireDB{
		Ref:   "/testRef",
		Title: "Test Request",
	}
	dbQuestions := []*gateway.QuestionDB{
		{
			Type:      1,
			Statement: "Как дела?",
			Order:     1,
		},
	}
	dbOptions := []*gateway.OptionsDB{
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

	coreQuestionnaire := dbQuestionnaire.ToCore(dbQuestions, dbOptions)

	assert.Equal(t, dbQuestionnaire.Ref, coreQuestionnaire.Ref)
	assert.Equal(t, dbQuestionnaire.Title, coreQuestionnaire.Title)
	assert.Equal(t, len(dbQuestions), len(coreQuestionnaire.Questions))
	assert.Equal(t, len(dbOptions), len(coreQuestionnaire.Questions[0].Options))
}

func TestMappingDbQuestionnaireFromCore(t *testing.T) {
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
	var dbQuestionnaire gateway.QuestionnaireDB
	dbQuestionnaire.FromCore(coreQuestionnaire)

	assert.Equal(t, dbQuestionnaire.Ref, coreQuestionnaire.Ref)
	assert.Equal(t, dbQuestionnaire.Title, coreQuestionnaire.Title)
}

func TestMappingDbQuestionToCore(t *testing.T) {
	dbQuestion := &gateway.QuestionDB{
		Statement: "Как дела?",
		Type:      0,
		Order:     1,
	}

	dbOptions := []*gateway.OptionsDB{
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

	coreQuestion := dbQuestion.ToCore(dbOptions)
	assert.Equal(t, dbQuestion.Statement, coreQuestion.Statement)
	assert.Equal(t, dbQuestion.Type, uint(coreQuestion.Type))
	assert.Equal(t, len(dbOptions), len(coreQuestion.Options))
}

func TestMappingDbQuestionFromCore(t *testing.T) {
	coreQuestion := &core.Question{
		Statement: "Как дела?",
		Type:      0,
		Options: []string{
			"Хорошо",
			"Нормально",
			"Плохо",
		},
	}
	var dbQuestion gateway.QuestionDB
	dbQuestion.FromCore(1, 1, coreQuestion)

	assert.Equal(t, dbQuestion.Statement, coreQuestion.Statement)
	assert.Equal(t, dbQuestion.Type, uint(coreQuestion.Type))
}
