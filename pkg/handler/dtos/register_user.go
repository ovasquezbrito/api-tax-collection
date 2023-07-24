package dtos

type RegisterUser struct {
	FirstLast string `json:"firstLast" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=8"`
	UriImg    string `json:"uri_img" validate:"required"`
}
