package controllers

import (
	models "Feedback-API/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

// GET /feedbacks
// Get all feedbacks
func FindFeedbacks(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var feedbacks []models.Feedback
	db.Find(&feedbacks)
	c.JSON(http.StatusOK, gin.H{"data": feedbacks})
}

// POST /feedbacks
// Create new feedbacks
func CreateFeedback(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	// Validate input
	var input models.CreateFeedbackInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Create Feedback
	feedback := models.Feedback{Name: input.Name, Email: input.Email, Ftype: input.Ftype, Text: input.Text}
	db.Create(&feedback)
	c.JSON(http.StatusOK, gin.H{"data": feedback})
}

// GET /feedbacks/:id
// Find a feedback
func FindFeedback(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var feedback models.Feedback
	if err := db.Where("id = ?", c.Param("id")).First(&feedback).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": feedback})
}

// DELETE /feedbacks/:id
// Delete a feedback
func DeleteFeedback(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var feedback models.Feedback
	if err := db.Where("id = ?", c.Param("id")).First(&feedback).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	db.Delete(&feedback)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
