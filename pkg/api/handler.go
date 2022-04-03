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

//
//func (h *RequestHandler) GetQuestionnaire(ctx context.Context, req *GetQuestionnaireRequest) (*GetQuestionnaireResponse, error) {
//	ref := req.GetRef()
//	Questionnaire, err := h.delegate_module.GetQuestionnaire(ref)
//	if err != nil {
//		return &GetQuestionnaireResponse{Questionnaire: nil}, err
//	}
//	return &GetQuestionnaireResponse{Questionnaire: Questionnaire}, nil
//}

func (h *RequestHandler) CreateQuestionnaire(ctx context.Context, req *CreateQuestionnaireRequest) (*CreateQuestionnaireResponse, error) {
	Questionnaire := req.GetQuestionnaire()
	ref, err := h.delegateModule.CreateQuestionnaire(Questionnaire)
	return &CreateQuestionnaireResponse{Ref: ref, Err: err.Error()}, nil
}

//
//func (h *RequestHandler) DeleteQuestionnaire(ctx context.Context, req *DeleteQuestionnaireRequest) (*DeleteQuestionnaireResponse, error) {
//	ref := req.GetRef()
//	result, err := h.delegate_module.DeleteQuestionnaire(ref)
//	if err != nil {
//		return &DeleteQuestionnaireResponse{Result: false}, err
//	}
//	return &DeleteQuestionnaireResponse{Result: result}, nil
//}
//
//func (h *RequestHandler) UpdateQuestionnaire(ctx context.Context, req *UpdateQuestionnaireRequest) (*UpdateQuestionnaireResponse, error) {
//	Questionnaire := req.GetQuestionnaire()
//	result, err := h.delegate_module.UpdateQuestionnaire(Questionnaire)
//	if err != nil {
//		return &UpdateQuestionnaireResponse{Result: false}, err
//	}
//	return &UpdateQuestionnaireResponse{Result: result}, nil
//}

func NewGrpcHandler(service *delegate.Module) (Handler, error) {
	return &RequestHandler{delegateModule: service}, nil
}
