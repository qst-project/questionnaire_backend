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

	//survey := &api.Survey{
	//	Id:    "2",
	//	Ref:   "/testRef",
	//	Title: "Test Request",
	//	Questions: []*api.Question{
	//		{Id: "1344", QuestionnaireId: "2", Text: "How are you?", Order: "1", Kind: "1"},
	//	},
	//	RadioPossibleAnswers: []*api.RadioPossibleAnswer{
	//		{Id: "143", Text: "Great", QuestionId: "1344"},
	//	},
	//	CheckboxPossibleAnswers: []*api.CheckboxPossibleAnswer{
	//		{Id: "14334", Text: "Great", QuestionId: "1344"},
	//	},
	//	TextPossibleAnswers: []*api.TextPossibleAnswer{
	//		{Id: "34341", Text: "Great", QuestionId: "1344", Placeholder: "Ответь здесь..."},
	//	},
	//	RadioAnswers: []*api.RadioAnswer{
	//		{Id: "1334", RadioPossibleAnswerId: "143"},
	//	},
	//	CheckboxAnswers: []*api.CheckboxAnswer{
	//		{Id: "1343", CheckboxPossibleAnswerId: "14334"},
	//	},
	//	TextAnswers: []*api.TextAnswer{
	//		{Id: "1343", TextPossibleAnswerId: "34341", Answer: "Bad"},
	//	},
	//}

	testRes, testResErr := c.Test(context.Background(), &api.TestRequest{})
	getSurveyRes, getSurveyErr := c.GetSurvey(context.Background(), &api.GetSurveyRequest{Ref: "TestRef"})
	//setSurveyRes, setSurveyResErr := c.CreateSurvey(context.Background(), &api.CreateSurveyRequest{Survey: survey})

	if testResErr != nil {
		log.Fatal(testResErr)
	}
	if getSurveyErr != nil {
		log.Fatal(getSurveyErr)
	}
	//if setSurveyRes != nil {
	//	log.Fatal(setSurveyResErr)
	//}
	log.Println(testRes.GetResult())
	log.Println(getSurveyRes.GetSurvey())
	//log.Println(setSurveyRes.GetResult())
}
