package main

import (
	"context"
	"github.com/qst-project/backend.git/pkg/api"
	"google.golang.org/grpc"
	"log"
)

// simple grpc_handler client simple to test server

func main() {
	conn, err := grpc.Dial("localhost:9091", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	c := api.NewQuestionnaireClient(conn)

	Questionnaire := &api.Questionnaire{
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
	//getQuestionnaireRes, getQuestionnaireErr := c.GetQuestionnaire(context.Background(), &api.GetQuestionnaireRequest{Ref: "/testQuestionnaire"})
	//createQuestionnaireRes, createQuestionnaireErr := c.CreateQuestionnaire(context.Background(), &api.CreateQuestionnaireRequest{QuestionnaireGateway: Questionnaire})
	//getQuestionnaireRes2, getQuestionnaireErr2 := c.GetQuestionnaire(context.Background(), &api.GetQuestionnaireRequest{Ref: "/testQuestionnaire"})
	//deleteQuestionnaireRes, deleteQuestionnaireErr := c.DeleteQuestionnaire(context.Background(), &api.DeleteQuestionnaireRequest{Ref: "/testQuestionnaire"})
	updateQuestionnaireRes, updateQuestionnaireErr := c.UpdateQuestionnaire(context.Background(), &api.UpdateQuestionnaireRequest{Questionnaire: Questionnaire})
	getQuestionnaireRes, getQuestionnaireErr := c.GetQuestionnaire(context.Background(), &api.GetQuestionnaireRequest{Ref: "/testQuestionnaire"})

	if testResErr != nil {
		log.Fatal(testResErr)
	}
	if getQuestionnaireErr != nil {
		log.Fatal(getQuestionnaireErr)
	}
	//if createQuestionnaireErr != nil {
	//	log.Fatal(createQuestionnaireErr)
	//}
	//if getQuestionnaireErr2 != nil {
	//	log.Fatal(getQuestionnaireErr2)
	//}
	//if deleteQuestionnaireErr != nil {
	//	log.Fatal(deleteQuestionnaireErr)
	//}
	if updateQuestionnaireErr != nil {
		log.Fatal(updateQuestionnaireErr)
	}
	log.Println(testRes.GetResult())
	log.Println(getQuestionnaireRes.GetQuestionnaire().Title)
	//log.Println(createQuestionnaireRes.GetResult())
	//log.Println(getQuestionnaireRes2.GetQuestionnaire())
	//log.Println(deleteQuestionnaireRes.GetResult())
	log.Println(updateQuestionnaireRes.GetResult())
}
