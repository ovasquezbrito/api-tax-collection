package entity

import (
	"time"
)

type RoleUser struct {
	Id           int       `db:"id"`
	NameOpcion   string    `db:"name_opcion" binding:"required"`
	Icon         string    `db:"icon" binding:"required"`
	ComponentUri string    `db:"componente_uri" binding:"required"`
	PageUrl      string    `db:"page_url" binding:"required"`
	OrderBy      int       `db:"orderby" binding:"required"`
	TypeOpcion   string    `db:"type_opcion" binding:"required"`
	NivelOpcion  int       `db:"nivel_opcion" binding:"required"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
	Status       bool      `db:"status"`
}
