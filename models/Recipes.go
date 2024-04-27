package models

import "gorm.io/gorm"

type Recipes struct {
	gorm.Model

	Title       string `gorm: "not null;unique_index`
	Description string
	Done        bool `gorm: "default:false"`
	UserID      uint
}
