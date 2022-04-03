package test

import (
	"context"
	"github.com/qst-project/backend.git/pkg/api"
	"google.golang.org/grpc"
	"log"
	"net"
	"testing"
)

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestGetQuestionnaireEndpoint(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := api.NewQuestionnaireClient(conn)
	resp, err := client.GetQuestionnaire(ctx, &api.GetQuestionnaireRequest{Ref: "TestRef"})
	if err != nil {
		t.Fatalf("Test failed: %v", err)
	}
	log.Printf("Response: %+v", resp.GetQuestionnaire())
}

func TestCreateQuestionnaireEndpoint(t *testing.T) {
	Questionnaire := &api.Questionnaire{
		Id:    "22",
		Ref:   "/testQuestionnaire",
		Title: "Test Title",
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

	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := api.NewQuestionnaireClient(conn)
	resp, err := client.CreateQuestionnaire(ctx, &api.CreateQuestionnaireRequest{Questionnaire: Questionnaire})
	if err != nil {
		t.Fatalf("Test failed: %v", err)
	}
	log.Printf("Response: %+v", resp.GetResult())
}

func TestUpdateQuestionnaireEndpoint(t *testing.T) {
	updateQuestionnaire := &api.Questionnaire{
		Id:    "22",
		Ref:   "/testQuestionnaire",
		Title: "Update Test Title",
		Questions: []*api.Question{
			{Id: "221", QuestionnaireId: "22", Text: "А Сколько ты зарабатываешь?", Order: "1", Kind: "1"},
			{Id: "222", QuestionnaireId: "22", Text: "Кто был президентом Беларуси?", Order: "2", Kind: "2"},
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

	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := api.NewQuestionnaireClient(conn)
	resp, err := client.UpdateQuestionnaire(ctx, &api.UpdateQuestionnaireRequest{Questionnaire: updateQuestionnaire})
	if err != nil {
		t.Fatalf("Test failed: %v", err)
	}
	log.Printf("Response: %+v", resp.GetResult())
}

func TestDeleteQuestionnaireEndpoint(t *testing.T) {

	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := api.NewQuestionnaireClient(conn)
	resp, err := client.DeleteQuestionnaire(ctx, &api.DeleteQuestionnaireRequest{Ref: "/testQuestionnaire"})
	if err != nil {
		t.Fatalf("Test failed: %v", err)
	}
	log.Printf("Response: %+v", resp.GetResult())
}
