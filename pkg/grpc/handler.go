package grpc

import (
	"context"
	"github.com/skinnykaen/quesionnaire_backend.git/pkg/api"
)

type RequestHandler struct {
	api.UnimplementedQuestionnaireServer
}

type Handler = api.QuestionnaireServer

func (h *RequestHandler) Test(ctx context.Context, req *api.TestRequest) (*api.TestResponse, error) {
	return &api.TestResponse{Result: "Test 1,2"}, nil
}
func NewGrpcHandler() (Handler, error) {
	return &RequestHandler{}, nil
}
