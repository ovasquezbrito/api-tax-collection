package models

import "strings"

type User struct {
	FirstLast string `json:"name" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" `
	UriImg    string `json:"uri_img" binding:"required"`
}

func (i User) UpperCase() *User {
	return &User{
		FirstLast: strings.ToUpper(i.FirstLast),
		Email:     strings.ToLower(i.Email),
		Password:  i.Password,
		UriImg:    i.UriImg,
	}
}

type LoginUserResponse struct {
	AccessToken string `json:"access_"`
	UserLogin   User   `json:"user"`
}
