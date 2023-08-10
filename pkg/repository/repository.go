package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/ovasquezbrito/tax-collection/pkg/entity"
)

type Authorization interface {
	CreateUser(ctx context.Context, user entity.User) (int, error)
	GetUser(ctx context.Context, email, password string) (entity.User, error)
	UpdateUser(ctx context.Context, userId int, user entity.User) (int, error)
	GetUserById(ctx context.Context, userId int) (*entity.UserResponse, error)
	GetUserByUserEmail(ctx context.Context, email string) (*entity.User, error)
	GetMenuOptionAll(ctx context.Context, IdUser int) ([]entity.RoleUser, error)
}

type RoleRepository interface {
	CreateRole(ctx context.Context, role entity.Role) (int, error)
	GetAll(ctx context.Context, query entity.QueryParameter) ([]entity.Role, int, error)
	GetById(ctx context.Context, idRol int) (*entity.Role, error)
	GetRoleByName(c context.Context, rolName string) (*entity.Role, error)
	GetUserByIdRole(ctx context.Context, idRole int) ([]entity.User, error)
	DeleteById(ctx context.Context, idRol int) (int64, error)
}
type UserRepository interface {
	GetAll(ctx context.Context, query entity.QueryParameter) ([]entity.UserResponse, int, error)
	AddRoleToUser(ctx context.Context, idRole, idUser int) error
}

type Repository struct {
	Authorization
	RoleRepository
	UserRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization:  NewAuthPostgres(db),
		RoleRepository: NewRolePostgres(db),
		UserRepository: NewUserPostgres(db),
	}
}
