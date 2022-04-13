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
	if err != nil {
		return &CreateQuestionnaireResponse{
			Ref: ref,
			Error: &Error{
				Code: 0,
				Text: err.Error(),
			},
		}, err
	}
	return &CreateQuestionnaireResponse{
		Ref: ref,
		Error: &Error{
			Code: 0,
			Text: "",
		},
	}, nil
}

func (h *RequestHandler) GetQuestionnaire(ctx context.Context, req *GetQuestionnaireRequest) (*GetQuestionnaireResponse, error) {
	questionnaire, err := h.QuestionnaireDelegate.GetQuestionnaire(req.GetRef())
	if err != nil {
		return &GetQuestionnaireResponse{
			Questionnaire: &questionnaire,
			Error: &Error{
				Code: 1,
				Text: err.Error(),
			},
		}, err
	}
	return &GetQuestionnaireResponse{
		Questionnaire: &questionnaire,
		Error: &Error{
			Code: 0,
			Text: "",
		},
	}, nil
}

func NewGrpcHandler(handlerArgs RequestHandlerArgs) Handler {
	return &RequestHandler{
		RequestHandlerArgs: handlerArgs,
	}
}
