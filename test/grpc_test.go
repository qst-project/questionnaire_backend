package test

import (
	"context"
	"github.com/qst-project/backend.git/pkg"
	"github.com/qst-project/backend.git/pkg/api"
	"github.com/qst-project/backend.git/pkg/configs"
	. "github.com/qst-project/backend.git/pkg/grpc"
	"github.com/qst-project/backend.git/pkg/repository"
	"github.com/qst-project/backend.git/pkg/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"log"
	"net"
	"testing"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
	lis = bufconn.Listen(bufSize)
	server := grpc.NewServer()
	logger := pkg.NewLogger()
	config, _ := configs.NewConfig()
	postgresClient, _ := repository.NewPostgresClient(config, logger)
	repos := repository.NewRepository(postgresClient)
	serv := service.NewService(repos)
	handler, _ := NewGrpcHandler(serv)
	api.RegisterQuestionnaireServer(server, handler)
	go func() {
		if err := server.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestFirstEndpoint(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := api.NewQuestionnaireClient(conn)
	resp, err := client.Test(ctx, &api.TestRequest{})
	if err != nil {
		t.Fatalf("Test failed: %v", err)
	}
	log.Printf("Response: %+v", resp)
	// Test for output here.
}

func TestGetSurveyEndpoint(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := api.NewQuestionnaireClient(conn)
	resp, err := client.GetSurvey(ctx, &api.GetSurveyRequest{Ref: "TestRef"})
	if err != nil {
		t.Fatalf("Test failed: %v", err)
	}
	log.Printf("Response: %+v", resp.GetSurvey())
}

func TestCreateSurveyEndpoint(t *testing.T) {
	survey := &api.Survey{
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
	resp, err := client.CreateSurvey(ctx, &api.CreateSurveyRequest{Survey: survey})
	if err != nil {
		t.Fatalf("Test failed: %v", err)
	}
	log.Printf("Response: %+v", resp.GetResult())
}

func TestUpdateSurveyEndpoint(t *testing.T) {
	updateSurvey := &api.Survey{
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
	resp, err := client.UpdateSurvey(ctx, &api.UpdateSurveyRequest{Survey: updateSurvey})
	if err != nil {
		t.Fatalf("Test failed: %v", err)
	}
	log.Printf("Response: %+v", resp.GetResult())
}

func TestDeleteSurveyEndpoint(t *testing.T) {

	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := api.NewQuestionnaireClient(conn)
	resp, err := client.DeleteSurvey(ctx, &api.DeleteSurveyRequest{Ref: "/testQuestionnaire"})
	if err != nil {
		t.Fatalf("Test failed: %v", err)
	}
	log.Printf("Response: %+v", resp.GetResult())
}
