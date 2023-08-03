package entity

import "strings"

type User struct {
	Id         int    `db:"id"`
	FirstLast  string `db:"first_last_name" binding:"required"`
	Email      string `db:"email" binding:"required"`
	AvatarUser string `db:"avatar_user" binding:"required"`
	IsAdmin    bool   `db:"is_admin"`
	FkRole     int    `db:"fk_role"`
	Password   string `db:"password" `
	Status     bool   `db:"status"`
	CreatedAt  string `db:"created_at"`
	UpdatedAt  string `db:"updated_at"`
}

func (i User) UpperCase() *User {
	return &User{
		FirstLast:  strings.ToUpper(i.FirstLast),
		Email:      strings.ToLower(i.Email),
		Password:   i.Password,
		AvatarUser: i.AvatarUser,
		IsAdmin:    i.IsAdmin,
		FkRole:     i.FkRole,
		Status:     i.Status,
		UpdatedAt:  i.UpdatedAt,
	}
}
