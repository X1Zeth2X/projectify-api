package models

import (
	"github.com/jinzhu/gorm"
)

// Project model
type Project struct {
	gorm.Model
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Images  []string `json:"images"`
}

// Add Create method
// Add Delete method
