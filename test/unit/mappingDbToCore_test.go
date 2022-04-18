package unit

import (
	"github.com/qst-project/backend.git/pkg/gateway"
	"github.com/qst-project/backend.git/test/unit/templates"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMappingDbQuestionnaireToCore(t *testing.T) {
	dbQuestionnaire := templates.GetDbQuestionnaire()
	dbQuestions := templates.GetDbQuestions()
	dbOptions := templates.GetDbOptions()

	coreQuestionnaire := dbQuestionnaire.ToCore(dbQuestions, dbOptions)

	assert.Equal(t, dbQuestionnaire.Ref, coreQuestionnaire.Ref)
	assert.Equal(t, dbQuestionnaire.Title, coreQuestionnaire.Title)
	assert.Equal(t, len(dbQuestions), len(coreQuestionnaire.Questions))
	assert.Equal(t, len(dbOptions), len(coreQuestionnaire.Questions[0].Options))
}

func TestMappingDbQuestionnaireFromCore(t *testing.T) {
	coreQuestionnaire := templates.GetCoreQuestionnaire()
	var dbQuestionnaire gateway.QuestionnaireDB
	dbQuestionnaire.FromCore(coreQuestionnaire)

	assert.Equal(t, dbQuestionnaire.Ref, coreQuestionnaire.Ref)
	assert.Equal(t, dbQuestionnaire.Title, coreQuestionnaire.Title)
}

func TestMappingDbQuestionToCore(t *testing.T) {
	dbQuestion := templates.GetDbQuestion()
	dbOptions := templates.GetDbOptions()

	coreQuestion := dbQuestion.ToCore(dbOptions)
	assert.Equal(t, dbQuestion.Statement, coreQuestion.Statement)
	assert.Equal(t, dbQuestion.Type, uint(coreQuestion.Type))
	assert.Equal(t, len(dbOptions), len(coreQuestion.Options))
}

func TestMappingDbQuestionFromCore(t *testing.T) {
	coreQuestion := templates.GetCoreQuestion()
	var dbQuestion gateway.QuestionDB
	dbQuestion.FromCore(1, 1, coreQuestion)

	assert.Equal(t, dbQuestion.Statement, coreQuestion.Statement)
	assert.Equal(t, dbQuestion.Type, uint(coreQuestion.Type))
}
