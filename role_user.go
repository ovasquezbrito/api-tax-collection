package baseapp

import (
	"time"
)

type RoleUser struct {
	Id           int       `json:"id" db:"id"`
	NameOpcion   string    `json:"name_opcion" db:"name_opcion" binding:"required"`
	Icon         string    `json:"icon" db:"icon" binding:"required"`
	ComponentUri string    `json:"component_page" db:"componente_uri" binding:"required"`
	PageUrl      string    `json:"page_url" db:"page_url" binding:"required"`
	OrderBy      int       `json:"order_by" db:"orderby" binding:"required"`
	TypeOpcion   string    `json:"type_opcion" db:"type_opcion" binding:"required"`
	NivelOpcion  int       `json:"nivel_opcion" db:"nivel_opcion" binding:"required"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
	Status       string    `json:"status" db:"status"`
}
