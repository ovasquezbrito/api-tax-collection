package models

import (
	"strings"
	"time"
)

type UserRole struct {
	UserID int64 `json:"user_id"`
	RoleID int64 `json:"role_id"`
}

type Role struct {
	IdRole    int       `json:"id_role"`
	RoleName  string    `json:"role_name" binding:"required"`
	RoleNivel int       `json:"role_nivel"`
	Status    bool      `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
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
	Status       bool      `json:"status" `
}

func (i Role) UpperCase() *Role {
	return &Role{
		RoleName:  strings.ToLower(i.RoleName),
		RoleNivel: i.RoleNivel,
	}
}
