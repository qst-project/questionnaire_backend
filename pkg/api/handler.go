package api

import (
	"context"
	"go.uber.org/fx"
)

type RequestHandlerArgs struct {
	fx.In

	QuestionnaireDelegate
}

type RequestHandler struct {
	UnimplementedQuestionnaireServiceServer
	RequestHandlerArgs
}

type Handler = QuestionnaireServiceServer

func (h *RequestHandler) CreateQuestionnaire(ctx context.Context, req *CreateQuestionnaireRequest) (*CreateQuestionnaireResponse, error) {
	ref, err := h.QuestionnaireDelegate.CreateQuestionnaire(req.GetQuestionnaire())
	return &CreateQuestionnaireResponse{
		Ref: ref,
		Error: &Error{
			Code:      0,
			TextError: err.Error(),
		},
	}, nil
}

func (h *RequestHandler) GetQuestionnaire(ctx context.Context, req *GetQuestionnaireRequest) (*GetQuestionnaireResponse, error) {
	questionnaire, err := h.QuestionnaireDelegate.GetQuestionnaire(req.GetRef())
	return &GetQuestionnaireResponse{
		Questionnaire: questionnaire,
		Error: &Error{
			Code:      0,
			TextError: err.Error(),
		},
	}, nil
}

func NewGrpcHandler(handlerArgs RequestHandlerArgs) Handler {
	return &RequestHandler{
		RequestHandlerArgs: handlerArgs,
	}
}
