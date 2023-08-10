package service

import (
	"context"

	"github.com/ovasquezbrito/tax-collection/pkg/models"
	"github.com/ovasquezbrito/tax-collection/pkg/repository"
	"github.com/ovasquezbrito/tax-collection/token"
	"github.com/ovasquezbrito/tax-collection/util"
)

type Authorization interface {
	CreateUser(ctx context.Context, user models.User) (int, error)
	LoginUser(ctx context.Context, email, password string) (models.LoginUserResponse, error)
	UpdateUser(ctx context.Context, idUser int, user models.User) (int, error)
	GetUserById(ctx context.Context, idUser int) (*models.UserResponse, error)
	ParseToken(token string) (string, error)
	VerifyToken(accessToken string) (*token.Payload, error)
	GetMenuOptionAll(ctx context.Context, IdUser int) ([]models.RoleUser, error)
}

type RoleService interface {
	CreateRole(ctx context.Context, user models.Role) (int, error)
	GetAll(ctx context.Context, query models.QueryParameter) ([]models.Role, int, error)
	GetById(ctx context.Context, idRol int) (*models.Role, error)
	DeleteById(ctx context.Context, idRol int) (int64, error)
}

type UsersServices interface {
	GetAll(ctx context.Context, query models.QueryParameter) ([]models.UserResponse, int, error)
	AddRoleToUser(ctx context.Context, userRole models.AsociateRoleToUser) error
}

type Service struct {
	Authorization
	RoleService
	UsersServices
}

func NewService(repos *repository.Repository, tokenMaker token.Maker, config util.Config) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization, tokenMaker, config),
		RoleService:   NewRoleService(repos.RoleRepository),
		UsersServices: NewUsersService(repos.UserRepository, repos.Authorization, repos.RoleRepository),
	}
}
