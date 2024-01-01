package model

import (
	"errors"
	"unicode"

	"gorm.io/gorm"
)

// Users
type User struct {
	gorm.Model
	Username string `json:"Username" gorm:"unique;not null;type:varchar(20)"`
	Password string `json:"Password" gorm:"not null;type:varchar(100);check:length(password)>=8"`

	// Relations
	ModeratedSubforums []Subforum `json:"Moderated_Subforums" gorm:"many2many:moderators_subforums"` // Many Users moderate many Subforums
	Posts []Post `json:"Posts" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"` // User makes many Posts
	Comments []Comment `json:"Comments" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"` // User makes many comments
	Votes []Vote `json:"Votes" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"` // Users make many votes

}

type Users struct {
	Users []User `json:"Users"`
}

func IsValidUser(u *User) error {
	containsSpace := func(s string) bool {
		for _, char := range s {
			if unicode.IsSpace(char) {
				return true
			}
		}
		return false
	}
	
	if len(u.Password) < 8 {
		return errors.New("password must be at least 8 characters long")
	}
	
	hasUpper := false
	hasLower := false
	for _, char := range u.Password {
		if unicode.IsUpper(char) {
			hasUpper = true
		} else if unicode.IsLower(char) {
			hasLower = true
		}
	}
	if !hasUpper || !hasLower {
		return errors.New("password must contain both uppercase and lowercase characters")
	}

	if containsSpace(u.Username) || containsSpace(u.Password) {
		return errors.New("username and password cannot contain spaces")
	}

	return nil
}