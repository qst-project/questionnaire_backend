package api

import (
	"context"
	"fmt"
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
		Ref:     ref,
		Err:     fmt.Sprintf("%v", err),
		ErrCode: 0,
	}, nil
}

func NewGrpcHandler(handlerArgs RequestHandlerArgs) Handler {
	return &RequestHandler{
		RequestHandlerArgs: handlerArgs,
	}
}
