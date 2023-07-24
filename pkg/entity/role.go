package entity

import (
	"strings"
	"time"
)

type Role struct {
	Id        int       `db:"id"`
	NameRole  string    `db:"name_role" binding:"required"`
	NivelRole int       `db:"nivel_role" binding:"required"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	Status    string    `db:"status"`
}

func (i Role) UpperCase() *Role {
	return &Role{
		NameRole:  strings.ToUpper(i.NameRole),
		NivelRole: i.NivelRole,
		Status:    i.Status,
	}
}
