package gateway

import (
	"github.com/qst-project/backend.git/pkg/core"
	"gorm.io/gorm"
	// "github.com/qst-project/backend.git/pkg/delegate"
)

type QuestionnaireGatewayImpl struct {
	postgresClient *PostgresClient
}

func (r *QuestionnaireGatewayImpl) GetQuestionnaire(ref string) (core.Questionnaire, error) {
	//TODO implement me
	panic("implement me")
}

func (r *QuestionnaireGatewayImpl) DeleteQuestionnaire(ref string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (r *QuestionnaireGatewayImpl) UpdateQuestionnaire(Questionnaire core.Questionnaire) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (r *QuestionnaireGatewayImpl) CreateQuestionnaire(Questionnaire core.Questionnaire) (id string, err error) {
	questionnaireDb := QuestionnaireDB{}
	questionnaireDb.FromCore(&Questionnaire)

	// TODO посоветоваться с владосом
	//var questionsDb []*QuestionDB
	//for order, coreQuestion := range Questionnaire.Questions {
	//	var dbQuestion QuestionDB
	//	dbQuestion.FromCore(uint(order), questionnaireDb.ID, coreQuestion)
	//	questionsDb = append(questionsDb, &dbQuestion)
	//}

	err = r.postgresClient.db.Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.Create(&questionnaireDb).Error; err != nil {
			return
		}
		return
	})
	return
}
