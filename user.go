package baseapp

import "strings"

type User struct {
	Id        int    `json:"id" db:"id"`
	FirstLast string `json:"name" db:"name" binding:"required"`
	Email     string `json:"email" db:"email" binding:"required"`
	Password  string `json:"password,omitempty" `
	UriImg    string `json:"uri_img" db:"uri_img" binding:"required"`
}

func (i User) UpperCase() *User {
	return &User{
		FirstLast: strings.ToUpper(i.FirstLast),
		Email:     strings.ToLower(i.Email),
		Password:  i.Password,
		UriImg:    i.UriImg,
	}
}
