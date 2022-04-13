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

func (r *QuestionnaireGatewayImpl) CreateQuestionnaire(Questionnaire core.Questionnaire) (ref string, err error) {
	questionnaireDb := QuestionnaireDB{}
	questionnaireDb.FromCore(&Questionnaire)

	err = r.postgresClient.db.Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.Create(&questionnaireDb).Error; err != nil {
			return
		}
		// TODO посоветоваться с владосом
		for order, coreQuestion := range Questionnaire.Questions {
			var questionDb QuestionDB
			questionDb.FromCore(uint(order), questionnaireDb.ID, coreQuestion)
			if err := tx.Create(&questionDb).Error; err != nil {
				return err
			}
			for _, coreOption := range coreQuestion.Options {
				var optionDb OptionsDB
				optionDb.FromCore(coreOption, uint(coreQuestion.Id))
				if err := tx.Create(&optionDb).Error; err != nil {
					return err
				}
			}
		}
		return nil
	})
	// TODO id
	return questionnaireDb.Ref, nil
}
