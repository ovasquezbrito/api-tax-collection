package dtos

type RegisterUser struct {
	FirstLast  string `json:"firstLast" validate:"required"`
	Email      string `json:"email" validate:"required,email"`
	Password   string `json:"password" validate:"required,min=8"`
	AvatarUser string `json:"avatar_user" validate:"required"`
}