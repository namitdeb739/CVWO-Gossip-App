package model

import (

	"gorm.io/gorm"
)

// Votes
type Vote struct {
	gorm.Model
	UserID uint `json:"UserID"`
	PostID uint `json:"PostID"`
	Type bool `json:"Type"`
}

type Votes struct {
	Votes []Vote `json:"Votes"`
}