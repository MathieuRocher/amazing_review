package main

import (
	"amazing_review/internal/adapter/application"
	"amazing_review/internal/adapter/handler"
	"amazing_review/internal/adapter/repository"
	"amazing_review/internal/infrastructure/database"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	router := gin.Default()

	// DATABASE
	database.InitDB()
	_ = database.DB.AutoMigrate(&repository.Review{}, &repository.FormAnswer{})

	// CACHE
	/*c := cache.New(10*time.Minute, 10*time.Minute)
	service.NewCacheService(c)*/

	// Importation des routes
	api := router.Group("/")

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

	reviewPort := os.Getenv("REVIEW_PORT")
	if reviewPort == "" {
		_ = godotenv.Load(".env") // charge localement si pas défini
		reviewPort = os.Getenv("REVIEW_PORT")
		if reviewPort == "" {
			reviewPort = "8082" // fallback si tout échoue
		}
	}

	err := router.Run("0.0.0.0:" + reviewPort)
	if err != nil {
		return
	}
}
