package models

import "strings"

type Result struct {
	Success bool        `json:"success,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type User struct {
	Email string
}

func NewUser(email string) *User {
	return &User{Email: email}
}

func (u *User) Name() string {
	data := strings.Split(u.Email, "@")
	if data != nil {
		return data[0]
	} else {
		return ""
	}
}
