package models

import "gorm.io/gorm"

type Recipes struct {
	gorm.Model

	Title        string `gorm:"not null;unique_index" json:"title"`
	Description  string `json:"description"`
	Instructions string `json:"instructions"`
	Ingredients  string `json:"ingredients"`
	Image        string `json:"image"`
	Done         bool   `gorm:"default:false" json:"done"`
	UserID       uint   `json:"user_id"`
}
