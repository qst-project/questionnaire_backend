package grpc

import (
	"context"
	// "fmt"
	"github.com/qst-project/backend.git/pkg/api"
	"github.com/qst-project/backend.git/pkg/service"
)

type RequestHandler struct {
	api.UnimplementedQuestionnaireServer
	service *service.Service
}

type Handler = api.QuestionnaireServer

func (h *RequestHandler) Test(ctx context.Context, req *api.TestRequest) (*api.TestResponse, error) {
	return &api.TestResponse{Result: "Test 1,2"}, nil
}

func (h *RequestHandler) GetSurvey(ctx context.Context, req *api.GetSurveyRequest) (*api.GetSurveyResponse, error) {
	ref := req.GetRef()
	survey, err := h.service.GetSurvey(ref)
	if err != nil {
		return &api.GetSurveyResponse{Survey: nil}, err
	}
	return &api.GetSurveyResponse{Survey: survey}, nil
}

func (h *RequestHandler) SetSurvey(ctx context.Context, req *api.CreateSurveyRequest) (*api.CreateSurveyResponse, error) {
	survey := req.GetSurvey()
	result, ref, err := h.service.SetSurvey(survey)
	if err != nil {
		return &api.CreateSurveyResponse{Result: false, Ref: ""}, err
	}
	return &api.CreateSurveyResponse{Result: result, Ref: ref}, nil
}

func (h *RequestHandler) DeleteSurvey(ctx context.Context, req *api.DeleteSurveyRequest) (*api.DeleteSurveyResponse, error) {
	ref := req.GetRef()
	result, err := h.service.DeleteSurvey(ref)
	if err != nil {
		return &api.DeleteSurveyResponse{Result: false}, err
	}
	return &api.DeleteSurveyResponse{Result: result}, nil
}

func NewGrpcHandler(service *service.Service) (Handler, error) {
	return &RequestHandler{service: service}, nil
}
