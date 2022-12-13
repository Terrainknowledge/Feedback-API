package models

type Feedback struct {
	ID    uint   `json:"id" gorm:"primary_key"`
	Name  string `json:"name" `
	Email string `json:"email"`
	Ftype string `json:"ftype"`
	Text  string `json:"text"`
}

type CreateFeedbackInput struct {
	Name  string `json:"name" binding:"required" `
	Email string `json:"email" binding:"required"`
	Ftype string `json:"ftype" binding:"required"`
	Text  string `json:"text" binding:"required"`
}
