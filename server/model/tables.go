package model

import "gorm.io/gorm"

// Users
type User struct {
	gorm.Model
	Username string `json:"Username" gorm:"primaryKey;unique;not null;type:varchar(20)"`
	Password string `json:"-" gorm:"not null;type:varchar(100)"`
	ModeratedSubforums []Subforum `json:"Moderated_Subforums" gorm:"many2many:user_subforums"`
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
	SubforumName string `json:"Name" gorm:"unique;not null;type:varchar(50)"`
	Moderators []User `json:"Moderators" gorm:"many2many:user_subforums"`
}

type Subforums struct {
	Subforums []Subforum `json:"Subforums"`
}

// Posts
type Post struct {
	gorm.Model
}

type Posts struct {
	Posts []Post `json:"Posts"`
}

// Commments
type Comment struct {
	gorm.Model
}

type Comments struct {
	Comments []Comment `json:"Comments"`
}

// Votes
type Vote struct {
	gorm.Model
}

type Votes struct {
	Votes []Vote `json:"Votes"`
}

// Tags
type Tag struct {
	gorm.Model
}

type Tags struct {
	Tags []Tag `json:"Tags"`
}