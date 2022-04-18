package api

import (
	"context"
	"github.com/qst-project/backend.git/pkg/delegate"
	"github.com/qst-project/backend.git/pkg/proto"
	"go.uber.org/fx"
)

type RequestHandlerArgs struct {
	fx.In

	delegate.QuestionnaireDelegate
}

type RequestHandler struct {
	proto.UnimplementedQuestionnaireServiceServer
	RequestHandlerArgs
}

type Handler = proto.QuestionnaireServiceServer

func (h *RequestHandler) CreateQuestionnaire(ctx context.Context, req *proto.CreateQuestionnaireRequest) (*proto.CreateQuestionnaireResponse, error) {
	ref, err := h.QuestionnaireDelegate.CreateQuestionnaire(req.GetQuestionnaire())
	if err != nil {
		return &proto.CreateQuestionnaireResponse{
			Ref: ref,
			Error: &proto.Error{
				Code: 0,
				Text: err.Error(),
			},
		}, err
	}
	return &proto.CreateQuestionnaireResponse{
		Ref: ref,
		Error: &proto.Error{
			Code: 0,
			Text: "",
		},
	}, nil
}

func (h *RequestHandler) GetQuestionnaire(ctx context.Context, req *proto.GetQuestionnaireRequest) (*proto.GetQuestionnaireResponse, error) {
	var protoQuestionnaire proto.Questionnaire
	qst, err := h.QuestionnaireDelegate.GetQuestionnaire(req.GetRef())
	protoQuestionnaire.FromCore(&qst)

	if err != nil {
		return &proto.GetQuestionnaireResponse{
			Questionnaire: &protoQuestionnaire,
			Error: &proto.Error{
				Code: 1,
				Text: err.Error(),
			},
		}, err
	}
	return &proto.GetQuestionnaireResponse{
		Questionnaire: &protoQuestionnaire,
		Error: &proto.Error{
			Code: 0,
			Text: "",
		},
	}, nil
}

func (h *RequestHandler) UpdateQuestionnaire(ctx context.Context, req *proto.UpdateQuestionnaireRequest) (*proto.UpdateQuestionnaireResponse, error) {
	err := h.QuestionnaireDelegate.UpdateQuestionnaire(req.GetQuestionnaire())
	if err != nil {
		return &proto.UpdateQuestionnaireResponse{
			Error: &proto.Error{
				Code: 0,
				Text: err.Error(),
			},
		}, err
	}
	return &proto.UpdateQuestionnaireResponse{
		Error: &proto.Error{
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
