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

type UserResponse struct {
	Id         int    `json:"uuid"`
	FirstLast  string `json:"first_last_name" `
	Email      string `json:"email" `
	AvatarUser string `json:"avatar_user"`
	IsAdmin    bool   `json:"is_admin"`
	FkRole     int    `json:"fk_role"`
	Status     bool   `json:"status"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
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

type AsociateRoleToUser struct {
	IdUser int `json:"id_user"`
	IdRole int `json:"id_role"`
}
