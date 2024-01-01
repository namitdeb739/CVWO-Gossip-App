package model

import (

	"gorm.io/gorm"
)

// Posts
type Post struct {
	gorm.Model
	UserID uint `json:"UserID"`
	SubforumID uint `json:"SubforumID"`
	Title string `json:"Title" gorm:"not null;type:varchar(50)"`
	Body string `json:"Body" gorm:"type:text"`

	// Relations
	Comments []Comment `json:"Comments" gorm:"foreignKey:PostID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"` // Posts have many Comments
	Votes []Vote `json:"Votes" gorm:"foreignKey:PostID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"` // Posts have many Votes
	Tags []Tag `json:"Tags" gorm:"many2many:posts_tags"` // Many Posts have many Tags
}

type Posts struct {
	Posts []Post `json:"Posts"`
}