package dtos

import (
	"time"

	"github.com/ovasquezbrito/tax-collection/pkg/models"
)

type Role struct {
	RoleName  string    `json:"role_name" binding:"required"`
	RoleNivel int       `json:"role_nivel" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Status    string    `json:"status"`
}

type RoleUser struct {
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

type GetAllRolesMenuResponse struct {
	Data []RoleUser `json:"data"`
}

type GetAllRolesResponse struct {
	Data  []models.Role `json:"data"`
	Total int           `json:"totalCount"`
}
