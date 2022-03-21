package grpc

import (
	"context"
	// "fmt"
	"github.com/skinnykaen/quesionnaire_backend.git/pkg/api"
	"github.com/skinnykaen/quesionnaire_backend.git/pkg/service"
)

type RequestHandler struct {
	api.UnimplementedQuestionnaireServer
	service *service.Service
}

type Handler = api.QuestionnaireServer

func (h *RequestHandler) Test(ctx context.Context, req *api.TestRequest) (*api.TestResponse, error) {
	return &api.TestResponse{Result: "Test 1,2"}, nil
}

func (h *RequestHandler) GetSurvey(ctx context.Context, req *api.SurveyRequest) (*api.SurveyResponse, error) {
	ref := req.GetRef()
	survey1, err := h.service.GetSurvey(ref)
	if err != nil{
		return nil, err
	} 
	// fmt.Printf(survey1)

	// survey := api.Survey {
	// 	Id: ref,
	// 	Title: "Test",
	// 	Questions: []*api.Survey_Question{
	// 		{Id: "1", Text: "Как дела?", Order: "1", Kind: "1"},
	// 		{Id: "2", Text: "Какой сегодня день?", Order: "2", Kind: "3"},
	// 	},
	// }
	return &api.SurveyResponse{Survey: &survey1}, nil
}

func NewGrpcHandler(service *service.Service) (Handler, error) {
	return &RequestHandler{service: service}, nil
}
