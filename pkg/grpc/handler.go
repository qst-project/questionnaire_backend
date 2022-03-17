package grpc

import (
	"context"
	"github.com/skinnykaen/quesionnaire_backend.git/pkg/api"
	// "github.com/skinnykaen/quesionnaire_backend.git/pkg/repository"
)

type RequestHandler struct {
	api.UnimplementedQuestionnaireServer
	// repository.NewPostgresClient
}

type Handler = api.QuestionnaireServer


func (h *RequestHandler) Test(ctx context.Context, req *api.TestRequest) (*api.TestResponse, error) {
	return &api.TestResponse{Result: "Test 1,2"}, nil
}

func (h *RequestHandler) GetSurvey(ctx context.Context, req *api.SurveyRequest) (*api.SurveyResponse, error) {
	// survey := repository.NewPostgresClient()
	survey := api.Survey {
		Id: "1",
		Title: "Test",
		Questions: []*api.Survey_Question{
			{Id: "1", Text: "Как дела?", Order: "1", Kind: "1"},
			{Id: "2", Text: "Какой сегодня день?", Order: "2", Kind: "3"},
		},
	}
	return &api.SurveyResponse{Survey: &survey}, nil
}

func NewGrpcHandler() (Handler, error) {
	return &RequestHandler{}, nil
}
