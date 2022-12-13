package main

import (
	"Feedback-API/controllers"
	"Feedback-API/models"
	"time"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	db := models.SetupModels()
	router.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, //add other sources ex: staging and production
		AllowMethods:     []string{"GET"},
		AllowHeaders:     []string{"Content-Type", "Content-Length", "Accept-Encoding", "Authorization", "Cache-Control"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	router.GET("/feedbacks", controllers.FindFeedbacks)
	router.POST("/feedbacks", controllers.CreateFeedback)
	router.GET("/feedbacks:id", controllers.FindFeedback)
	router.DELETE("/feedbacks:id", controllers.DeleteFeedback)

	router.Run("localhost:8090")
}
