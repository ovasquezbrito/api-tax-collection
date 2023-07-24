package repository

import (
	"github.com/jmoiron/sqlx"
	baseapp "github.com/ovasquezbrito/base-app"
)

type Authorization interface {
	CreateUser(user baseapp.User) (int, error)
	GetUser(email, password string) (baseapp.User, error)
	UpdateUser(userId int, user baseapp.User) (int, error)
	GetUserById(userId int) (baseapp.User, error)
	GetUserByUserName(email string) (int, error)
	GetUserByUserEmail(email string) (baseapp.User, error)
	GetMenuOptionAll(IdUser int) ([]baseapp.RoleUser, error)
}

type RoleRepository interface {
	GetAll(query baseapp.QueryParameter) ([]baseapp.Role, int, error)
	GetById(idRol int) (baseapp.Role, error)
}

type Repository struct {
	Authorization
	RoleRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization:  NewAuthPostgres(db),
		RoleRepository: NewRolePostgres(db),
	}
}
