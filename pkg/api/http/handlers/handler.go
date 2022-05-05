package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/qst-project/backend.git/pkg/delegate"
)

type RequestHandler struct {
	//fx.In
	delegate.HttpQuestionnaireDelegate
}

func NewHandler(httpQuestionnaireDelegate delegate.HttpQuestionnaireDelegate) RequestHandler {
	return RequestHandler{
		HttpQuestionnaireDelegate: httpQuestionnaireDelegate,
	}
}

func (h *RequestHandler) InitRoutes() *gin.Engine {
	router := gin.New()
	questionaire := router.Group("/questionnaire")
	{
		questionaire.GET("/", h.GetQuestionnaire)
		questionaire.POST("/create", h.CreateQuestionnaire)
		questionaire.PUT("/update", h.UpdateQuestionnaire)
	}
	return router
}
