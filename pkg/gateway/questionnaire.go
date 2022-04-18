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
	var questionnaireDb QuestionnaireDB
	var questionsDb []*QuestionDB
	var optionsDb []*OptionsDB
	err := r.postgresClient.db.Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.Where("ref = ?", ref).First(&questionnaireDb).Error; err != nil {
			return
		}
		if err = tx.Where("questionnaire_id = ?", &questionnaireDb.ID).Find(&questionsDb).Error; err != nil {
			return
		}
		for _, questionDb := range questionsDb {
			var tempOptionsDb []*OptionsDB
			if err = tx.Where("question_id = ?", &questionDb.ID).Find(&tempOptionsDb).Error; err != nil {
				return
			}
			optionsDb = append(optionsDb, tempOptionsDb...)
		}
		return
	})
	coreQuestionnaire := questionnaireDb.ToCore(questionsDb, optionsDb)
	if err != nil {
		return coreQuestionnaire, err
	}
	return coreQuestionnaire, nil
}

func (r *QuestionnaireGatewayImpl) DeleteQuestionnaire(ref string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (r *QuestionnaireGatewayImpl) UpdateQuestionnaire(Questionnaire core.Questionnaire) error {
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
	if err != nil {
		return "", err
	}
	// TODO gen ref
	return questionnaireDb.Ref, nil
}
