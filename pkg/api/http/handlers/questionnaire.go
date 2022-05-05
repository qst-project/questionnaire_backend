package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/qst-project/backend.git/pkg/httpModels"
	"net/http"
)

type getResponse struct {
	questionnaire httpModels.Questionnaire `json:"questionnaire"`
}

func (h *RequestHandler) GetQuestionnaire(c *gin.Context) {
	fmt.Println("Get")
	//ref := c.Param("ref")
	//fmt.Println(ref)
	questionnaire, err := h.HttpQuestionnaireDelegate.GetQuestionnaire("testRef")
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, &getResponse{
		questionnaire: questionnaire,
	})
}

type createInput struct {
	Questionnaire httpModels.Questionnaire `json:"questionnaire"`
}

type createResponse struct {
	Ref string `json:"ref"`
}

func (h *RequestHandler) CreateQuestionnaire(c *gin.Context) {
	fmt.Println("Create")
	inp := new(createInput)
	if err := c.BindJSON(inp); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	ref, err := h.HttpQuestionnaireDelegate.CreateQuestionnaire(&inp.Questionnaire)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, &createResponse{
		Ref: ref,
	})
}

func (h *RequestHandler) UpdateQuestionnaire(c *gin.Context) {
	fmt.Println("Update")
	c.Status(http.StatusOK)
}
