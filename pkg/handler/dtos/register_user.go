package dtos

type RegisterUser struct {
	FirstLast  string `json:"first_last_name" validate:"required" swaggertype:"string" example:"usuario administrador"`
	Email      string `json:"email" validate:"required,email" swaggertype:"string" example:"admin@gmail.com"`
	Password   string `json:"password" validate:"required,min=8" swaggertype:"string" example:"1234567890"`
	AvatarUser string `json:"avatar_user" validate:"required" swaggertype:"string" format:"base64" example:"U3dhZ2dlciByb2Nrcw=="`
}
