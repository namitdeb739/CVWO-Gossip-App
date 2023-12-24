package model

import (
	"gorm.io/gorm"
)

// Users
type User struct {
	gorm.Model
	Username string `json:"Username" gorm:"unique;not null;type:varchar(20)"`
	Password string `json:"password" gorm:"not null;type:varchar(100);check:length(password)>=8"`

	// Relations
	ModeratedSubforums []Subforum `json:"Moderated_Subforums" gorm:"many2many:moderators_subforums"` // Many Users moderate many Subforums
	Posts []Post `json:"Posts" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"` // User makes many Posts
	Comments []Comment `json:"Comments" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"` // User makes many comments
	Votes []Vote `json:"Votes" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"` // Users make many votes
}

type Users struct {
	Users []User `json:"Users"`
}

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