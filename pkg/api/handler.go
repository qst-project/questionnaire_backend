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
		Ref:     ref,
		Err:     err.Error(),
		ErrCode: 0,
	}, nil
}

func NewGrpcHandler(RequestHandlerArgs) (Handler, error) {
	return &RequestHandler{
		UnimplementedQuestionnaireServiceServer: UnimplementedQuestionnaireServiceServer{},
	}, nil
}
