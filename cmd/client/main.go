package main

import (
	"context"
	"github.com/qst-project/backend.git/pkg/api"
	"google.golang.org/grpc"
	"log"
)

// simple grpc client simple to test server

func main() {
	conn, err := grpc.Dial("localhost:9091", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	c := api.NewQuestionnaireClient(conn)

	survey := &api.Survey{
		Id:    "22",
		Ref:   "/testQuestionnaire",
		Title: "Update Test Title",
		Questions: []*api.Question{
			{Id: "221", QuestionnaireId: "22", Text: "Сколько ты зарабатываешь?", Order: "1", Kind: "1"},
			{Id: "222", QuestionnaireId: "22", Text: "Кто был президентом России?", Order: "2", Kind: "2"},
		},
		RadioPossibleAnswers: []*api.RadioPossibleAnswer{
			{Id: "221", Text: ">100.000", QuestionId: "221"},
			{Id: "222", Text: "<100.000", QuestionId: "221"},
		},
		CheckboxPossibleAnswers: []*api.CheckboxPossibleAnswer{
			{Id: "221", Text: "Путин", QuestionId: "222"},
			{Id: "222", Text: "Лукашенко", QuestionId: "222"},
			{Id: "223", Text: "Ельцин", QuestionId: "222"},
		},
	}

	testRes, testResErr := c.Test(context.Background(), &api.TestRequest{})
	//getSurveyRes, getSurveyErr := c.GetSurvey(context.Background(), &api.GetSurveyRequest{Ref: "/testQuestionnaire"})
	//createSurveyRes, createSurveyErr := c.CreateSurvey(context.Background(), &api.CreateSurveyRequest{Survey: survey})
	//getSurveyRes2, getSurveyErr2 := c.GetSurvey(context.Background(), &api.GetSurveyRequest{Ref: "/testQuestionnaire"})
	//deleteSurveyRes, deleteSurveyErr := c.DeleteSurvey(context.Background(), &api.DeleteSurveyRequest{Ref: "/testQuestionnaire"})
	updateSurveyRes, updateSurveyErr := c.UpdateSurvey(context.Background(), &api.UpdateSurveyRequest{Survey: survey})
	getSurveyRes, getSurveyErr := c.GetSurvey(context.Background(), &api.GetSurveyRequest{Ref: "/testQuestionnaire"})

	if testResErr != nil {
		log.Fatal(testResErr)
	}
	if getSurveyErr != nil {
		log.Fatal(getSurveyErr)
	}
	//if createSurveyErr != nil {
	//	log.Fatal(createSurveyErr)
	//}
	//if getSurveyErr2 != nil {
	//	log.Fatal(getSurveyErr2)
	//}
	//if deleteSurveyErr != nil {
	//	log.Fatal(deleteSurveyErr)
	//}
	if updateSurveyErr != nil {
		log.Fatal(updateSurveyErr)
	}
	log.Println(testRes.GetResult())
	log.Println(getSurveyRes.GetSurvey().Title)
	//log.Println(createSurveyRes.GetResult())
	//log.Println(getSurveyRes2.GetSurvey())
	//log.Println(deleteSurveyRes.GetResult())
	log.Println(updateSurveyRes.GetResult())
}
