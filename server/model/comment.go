package model

import (

	"gorm.io/gorm"
)

// Commments
type Comment struct {
	gorm.Model
	UserID uint `json:"UserID"`
	PostID uint `json:"PostID"`
	ParentCommentID *uint `json:"ParentCommentID"`

	// Relations
	ChildrenComments []Comment `json:"ChildrenComments" gorm:"foreignKey:ParentCommentID"` // Comments have many children Comments
}

type Comments struct {
	Comments []Comment `json:"Comments"`
}