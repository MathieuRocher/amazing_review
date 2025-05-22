package handler

import (
	"amazing_review/internal/adapter/application"
	"amazing_review/internal/adapter/handler/dto/review"
	"net/http"
	"strconv"

	domain "github.com/MathieuRocher/amazing_domain"
	"github.com/gin-gonic/gin"
)

type ReviewHandler struct {
	useCase application.ReviewUseCaseInterface
}

func NewReviewHandler(uc application.ReviewUseCaseInterface) *ReviewHandler {
	return &ReviewHandler{useCase: uc}
}

func (h *ReviewHandler) GetReviews(c *gin.Context) {
	// Récupération des query params
	pageStr := c.Query("page")
	limitStr := c.Query("limit")
	var (
		reviews []domain.Review
		err     error
	)

	if pageStr != "" && limitStr != "" {
		page, err1 := strconv.Atoi(pageStr)
		limit, err2 := strconv.Atoi(limitStr)

		if err1 == nil && err2 == nil {
			// Appel avec pagination
			reviews, err = h.useCase.FindAllWithPagination(page, limit)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
		}
	} else {
		reviews, err = h.useCase.FindAll()
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var out []review.ReviewOutput
	for _, r := range reviews {
		out = append(out, *review.FromDomain(&r))
	}

	c.JSON(http.StatusOK, gin.H{
		"reviews": out,
	})
}

func (h *ReviewHandler) CreateReview(c *gin.Context) {
	var input review.ReviewInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input", "details": err.Error()})
		return
	}

	uidRaw, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	userID, ok := uidRaw.(uint)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id in context"})
		return
	}

	if err := h.useCase.Create(&input, userID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "review created successfully"})
}

func (h *ReviewHandler) GetReviewByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	r, err := h.useCase.FindByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "review not found"})
		return
	}

	output := review.FromDomain(r)
	c.JSON(http.StatusOK, gin.H{"review": output})
}

func (h *ReviewHandler) UpdateReviewByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	var input review.ReviewInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	existingReview, err := h.useCase.FindByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "review not found"})
		return
	}

	existingReview.FormID = input.FormId
	existingReview.CourseAssignmentID = input.CourseAssignmentId
	existingReview.FormAnswers = nil // gestion des answers ignorée ici

	if err := h.useCase.Update(existingReview); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update review"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "review updated successfully"})
}

func (h *ReviewHandler) DeleteReviewByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	if err := h.useCase.Delete(uint(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "review not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "review deleted successfully"})
}

func (h *ReviewHandler) RegisterRoutes(rg *gin.RouterGroup) {
	group := rg.Group("/reviews")
	{
		group.GET("", h.GetReviews)
		group.GET(":id", h.GetReviewByID)
		group.POST("", h.CreateReview)
		group.PUT(":id", h.UpdateReviewByID)
		group.DELETE(":id", h.DeleteReviewByID)
	}
}
