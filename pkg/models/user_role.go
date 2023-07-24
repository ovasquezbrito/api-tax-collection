package models

import "time"

type UserRole struct {
	UserID int64 `json:"user_id"`
	RoleID int64 `json:"role_id"`
}

type Role struct {
	IdRole    int       `json:"id_role"`
	NameRole  string    `json:"name_role" binding:"required"`
	NivelRole int       `json:"nivel_role" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Status    string    `json:"status"`
}

type RoleUser struct {
	IdRole       int       `json:"id_role"`
	NameOpcion   string    `json:"name_opcion" binding:"required"`
	Icon         string    `json:"icon" binding:"required"`
	ComponentUri string    `json:"component_page" binding:"required"`
	PageUrl      string    `json:"page_url" binding:"required"`
	OrderBy      int       `json:"order_by" binding:"required"`
	TypeOpcion   string    `json:"type_opcion" binding:"required"`
	NivelOpcion  int       `json:"nivel_opcion" binding:"required"`
	CreatedAt    time.Time `json:"created_at" `
	UpdatedAt    time.Time `json:"updated_at" `
	Status       string    `json:"status" `
}
