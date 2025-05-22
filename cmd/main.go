package main

import (
	"amazing_review/internal/adapter/application"
	"amazing_review/internal/adapter/handler"
	"amazing_review/internal/adapter/repository"
	"amazing_review/internal/infrastructure/database"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	// DATABASE
	database.InitDB()

	// CACHE
	/*c := cache.New(10*time.Minute, 10*time.Minute)
	service.NewCacheService(c)*/

	// Importation des routes
	api := router.Group("/api")
	api.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello world",
		})
	})

	// FORM ANSWERS
	formAnswersRepository := repository.NewFormAnswerRepository()
	formAnswersUseCase := application.NewFormAnswerUseCase(formAnswersRepository)
	formAnswersHandler := handler.NewFormAnswerHandler(formAnswersUseCase)
	formAnswersHandler.RegisterRoutes(api)

	// REVIEWS
	reviewRepository := repository.NewReviewRepository()
	reviewUseCase := application.NewReviewUseCase(
		reviewRepository,
		formAnswersRepository,
	)
	reviewHandler := handler.NewReviewHandler(reviewUseCase)
	reviewHandler.RegisterRoutes(api)

	err := router.Run("localhost:8082")
	if err != nil {
		return
	}
}
