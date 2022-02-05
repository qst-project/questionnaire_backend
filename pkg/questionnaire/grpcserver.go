package questionnaire

import (
	"context"
	"github.com/skinnykaen/quesionnaire_backend.git/pkg/api"
)

type handler struct {
	api.UnimplementedQuestionnaireServer
}

type Handler = api.QuestionnaireServer

func (h *handler) Test(ctx context.Context, req *api.TestRequest) (*api.TestResponse, error){
	return & api.TestResponse{Result: "Test 1,2"}, nil
}
func New()(Handler, error) {
	return &handler{}, nil
}