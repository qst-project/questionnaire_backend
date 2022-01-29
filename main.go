package main

import (
	"github.com/skinnykaen/quesionnaire_backend.git/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func main(){
	dsn := "host=localhost user=skinny password=12345 dbname=questionnaire_backend port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&models.User{},
				   &models.Question{},
				   &models.Questionnaire{},
				   &models.RadioPossibleAnswer{},
				   &models.TextPossibleAnswer{},
				   &models.CheckboxPossibleAnswer{},
				   &models.CheckboxAnswer{},
				   &models.RadioAnswer{},
				   &models.TextAnswer{},
				   )
	db.Create(&models.User{GoogleId: "1test"})
}