package dtos

type RegisterUser struct {
	FirstLast  string `json:"first_last_name" validate:"required"`
	Email      string `json:"email" validate:"required,email"`
	Password   string `json:"password" validate:"required,min=8"`
	FkRole     int    `json:"fk_role"`
	AvatarUser string `json:"avatar_user" validate:"required"`
}
