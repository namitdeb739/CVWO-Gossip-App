package model

import (

	"gorm.io/gorm"
)

// Subforums
type Subforum struct {
	gorm.Model
	Name string `json:"Name" gorm:"unique;not null;type:varchar(50)"`
	Description string `json:"Description" gorm:"type:text"`

	// Relations
	Moderators []User `json:"Moderators" gorm:"many2many:moderators_subforums"` // Many Subforums are moderated by many Users
	Posts []Post `json:"Posts" gorm:"foreignKey:SubforumID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"` // Subforum has many Posts
}

type Subforums struct {
	Subforums []Subforum `json:"Subforums"`
}