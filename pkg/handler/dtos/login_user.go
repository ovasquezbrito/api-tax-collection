package dtos

type LoginUser struct {
	Email    string `json:"email" validate:"required,email" swaggertype:"string" format:"string" example:"admin@gmail.com"`
	Password string `json:"password" validate:"required" swaggertype:"string" format:"string" example:"1234567890"`
}

type UserResponse struct {
	ID         int    `json:"id"`
	Email      string `json:"email"`
	Name       string `json:"name"`
	AvatarUser string `json:"avatar_user"`
	Token      string `json:"token"`
}
