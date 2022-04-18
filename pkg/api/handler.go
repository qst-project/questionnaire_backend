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
	var protoQuestionnaire Questionnaire
	qst, err := h.QuestionnaireDelegate.GetQuestionnaire(req.GetRef())
	protoQuestionnaire.FromCore(&qst)

	if err != nil {
		return &GetQuestionnaireResponse{
			Questionnaire: &protoQuestionnaire,
			Error: &Error{
				Code: 1,
				Text: err.Error(),
			},
		}, err
	}
	return &GetQuestionnaireResponse{
		Questionnaire: &protoQuestionnaire,
		Error: &Error{
			Code: 0,
			Text: "",
		},
	}, nil
}

func (h *RequestHandler) UpdateQuestionnaire(ctx context.Context, req *UpdateQuestionnaireRequest) (*UpdateQuestionnaireResponse, error) {
	err := h.QuestionnaireDelegate.UpdateQuestionnaire(req.GetQuestionnaire())
	if err != nil {
		return &UpdateQuestionnaireResponse{
			Error: &Error{
				Code: 0,
				Text: err.Error(),
			},
		}, err
	}
	return &UpdateQuestionnaireResponse{
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
