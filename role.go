package baseapp

import (
	"strings"
	"time"
)

type Role struct {
	Id        int       `json:"id" db:"id"`
	NameRole  string    `json:"name_role" db:"name_role" binding:"required"`
	NivelRole int       `json:"nivel_role" db:"nivel_role" binding:"required"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	Status    string    `json:"status" db:"status"`
}

func (i Role) UpperCase() *Role {
	return &Role{
		NameRole:  strings.ToUpper(i.NameRole),
		NivelRole: i.NivelRole,
		Status:    i.Status,
	}
}
