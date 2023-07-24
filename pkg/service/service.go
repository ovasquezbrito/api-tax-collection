package service

import (
	baseapp "github.com/ovasquezbrito/base-app"
	"github.com/ovasquezbrito/base-app/pkg/repository"
	"github.com/ovasquezbrito/base-app/token"
	"github.com/ovasquezbrito/base-app/util"
)

type Authorization interface {
	CreateUser(user baseapp.User) (int, error)
	LoginUser(email, password string) (loginUserResponse, error)
	UpdateUser(idUser int, user baseapp.User) (int, error)
	GetUserById(idUser int) (baseapp.User, error)
	GetUserByUserName(email string) (int, error)
	ParseToken(token string) (string, error)
	VerifyToken(accessToken string) (*token.Payload, error)
	GetMenuOptionAll(IdUser int) ([]baseapp.RoleUser, error)
}

type RoleService interface {
	GetAll(query baseapp.QueryParameter) ([]baseapp.Role, int, error)
	GetById(idRol int) (baseapp.Role, error)
}

type Service struct {
	Authorization
	RoleService
}

func NewService(repos *repository.Repository, tokenMaker token.Maker, config util.Config) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization, tokenMaker, config),
		RoleService:   NewRoleService(repos.RoleRepository),
	}
}
