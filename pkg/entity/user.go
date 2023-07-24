package entity

import "strings"

type User struct {
	Id        int    `db:"id"`
	FirstLast string `db:"name" binding:"required"`
	Email     string `db:"email" binding:"required"`
	Password  string `db:"password" `
	UriImg    string `db:"uri_img" binding:"required"`
}

func (i User) UpperCase() *User {
	return &User{
		FirstLast: strings.ToUpper(i.FirstLast),
		Email:     strings.ToLower(i.Email),
		Password:  i.Password,
		UriImg:    i.UriImg,
	}
}
