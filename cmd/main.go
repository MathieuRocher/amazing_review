package main

import (
	"amazing_review/internal/adapter/application"
	"amazing_review/internal/adapter/handler"
	"amazing_review/internal/adapter/repository"
	"amazing_review/internal/infrastructure/database"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	router := gin.Default()

	e := godotenv.Load()
	if e != nil {
		log.Fatal("Error loading .env file")
	}

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
		log.Fatal("REVIEW_PORT must be set in environment")
	}

	err := router.Run("localhost:" + reviewPort)
	if err != nil {
		return
	}
}
