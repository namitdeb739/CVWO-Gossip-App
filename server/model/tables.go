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
	ModeratedSubforums []Subforum `json:"Moderated_Subforums" gorm:"many2many:moderators_subforums"`
	Posts []Post `json:"Posts" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Comments []Comment `json:"Comments" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Votes []Vote `json:"Votes" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
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
	Moderators []User `json:"Moderators" gorm:"many2many:moderators_subforums"`
	Posts []Post `json:"Posts" gorm:"foreignKey:SubforumID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
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
	Comments []Comment `json:"Comments" gorm:"foreignKey:PostID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Votes []Vote `json:"Votes" gorm:"foreignKey:PostID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Tags []Tag `json:"Tags" gorm:"many2many:posts_tags"`
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
	ChildrenComments []Comment `json:"ChildrenComments" gorm:"foreignKey:ParentCommentID"`
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
	Posts []Post `json:"Posts" gorm:"many2many:posts_tags"`
}

type Tags struct {
	Tags []Tag `json:"Tags"`
}