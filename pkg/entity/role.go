package entity

import (
	"strings"
	"time"
)

type Role struct {
	Id        int       `db:"id"`
	RoleName  string    `db:"role_name" binding:"required"`
	RoleNivel int       `db:"role_nivel"`
	Status    bool      `db:"status"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func (i Role) UpperCase() *Role {
	return &Role{
		RoleName:  strings.ToLower(i.RoleName),
		RoleNivel: i.RoleNivel,
		Status:    i.Status,
	}
}
