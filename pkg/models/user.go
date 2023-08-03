package models

import "strings"

type User struct {
	Id         int    `json:"id"`
	FirstLast  string `json:"first_last_name" binding:"required"`
	Email      string `json:"email" binding:"required"`
	Password   string `json:"password" `
	AvatarUser string `json:"avatar_user" binding:"required"`
	IsAdmin    bool   `json:"is_admin"`
	FkRole     int    `json:"fk_role"`
}

func (i User) UpperCase() *User {
	return &User{
		FirstLast:  strings.ToUpper(i.FirstLast),
		Email:      strings.ToLower(i.Email),
		Password:   i.Password,
		AvatarUser: i.AvatarUser,
	}
}

type LoginUserResponse struct {
	AccessToken string `json:"access_"`
	UserLogin   *User  `json:"user"`
}
