package api

import (
	"context"
	"github.com/qst-project/backend.git/pkg/delegate"
)

type RequestHandler struct {
	UnimplementedQuestionnaireServiceServer

	delegateModule *delegate.Module
}

type Handler = QuestionnaireServiceServer

func (h *RequestHandler) CreateQuestionnaire(ctx context.Context, req *CreateQuestionnaireRequest) (*CreateQuestionnaireResponse, error) {
	ref, err := h.delegateModule.CreateQuestionnaire(req.GetQuestionnaire())
	return &CreateQuestionnaireResponse{
		Ref:     ref,
		Err:     err.Error(),
		ErrCode: 0,
	}, nil
}

func NewGrpcHandler(service *delegate.Module) (Handler, error) {
	return &RequestHandler{delegateModule: service}, nil
}
