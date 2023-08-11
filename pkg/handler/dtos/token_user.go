package dtos

type TokenUser struct {
	Token string `json:"token" binding:"required"`
}

type TokenResponse struct {
	Token string `json:"token"`
}
