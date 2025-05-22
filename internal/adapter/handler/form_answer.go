package handler

import (
	"amazing_review/internal/adapter/application"
	domain "github.com/MathieuRocher/amazing_domain"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type FormAnswerHandler struct {
	useCase application.FormAnswerUseCaseInterface
}

func NewFormAnswerHandler(uc application.FormAnswerUseCaseInterface) *FormAnswerHandler {
	return &FormAnswerHandler{useCase: uc}
}

func (h *FormAnswerHandler) GetFormAnswers(c *gin.Context) {
	formAnswers, _ := h.useCase.FindAll()

	c.JSON(http.StatusOK, gin.H{
		"message": formAnswers,
	})
}

func (h *FormAnswerHandler) CreateFormAnswer(c *gin.Context) {
	var payload domain.FormAnswer
	err := c.Bind(&payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	err = h.useCase.Create(&payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "unknow sql error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "formAnswer created"})
	return
}

func (h *FormAnswerHandler) GetFormAnswerByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	formAnswer, err := h.useCase.FindByID(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "formAnswer not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"formAnswer": formAnswer})
	return
}

func (h *FormAnswerHandler) UpdateFormAnswerByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	// est-ce qu'il existe
	formAnswer, err := h.useCase.FindByID(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "formAnswer not found"})
		return
	}

	// je le remplace par le body
	err = c.Bind(&formAnswer)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	// je save
	err = h.useCase.Update(formAnswer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "formAnswer Updated"})
	return
}

func (h *FormAnswerHandler) DeleteFormAnswerByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := h.useCase.Delete(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "formAnswer not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "formAnswer deleted succesfully"})
	return

}

func (h *FormAnswerHandler) RegisterRoutes(rg *gin.RouterGroup) {
	group := rg.Group("/form-answers")
	{
		group.GET("", h.GetFormAnswers)
		group.POST("", h.CreateFormAnswer)
		group.GET(":id", h.GetFormAnswerByID)
		group.PUT(":id", h.UpdateFormAnswerByID)
		group.DELETE(":id", h.DeleteFormAnswerByID)
	}
}
