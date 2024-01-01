package model

import (

	"gorm.io/gorm"
)

// Tags
type Tag struct {
	gorm.Model
	Name string `json:"Name" gorm:"unique;not null;type:varchar(20)"`

	// Relations
	Posts []Post `json:"Posts" gorm:"many2many:posts_tags"` // Many Tags have many Posts
}

type Tags struct {
	Tags []Tag `json:"Tags"`
}