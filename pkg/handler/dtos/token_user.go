package dtos

type TokenUser struct {
	Token string `json:"token" binding:"required"`
}
