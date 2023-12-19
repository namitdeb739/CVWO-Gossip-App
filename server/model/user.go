package model

import (
	"github.com/google/uuid"
 	"gorm.io/gorm"
)

type User struct {
	User_ID string `json:"User_ID"`
}

type Users struct {
	Users []User `json:"Users"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.User_ID = uuid.New()
	return
}