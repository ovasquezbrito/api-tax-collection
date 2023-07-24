package service

import (
	"github.com/ovasquezbrito/tax-collection/pkg/models"
	"github.com/ovasquezbrito/tax-collection/pkg/repository"
	"github.com/ovasquezbrito/tax-collection/token"
	"github.com/ovasquezbrito/tax-collection/util"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	LoginUser(email, password string) (models.LoginUserResponse, error)
	UpdateUser(idUser int, user models.User) (int, error)
	GetUserById(idUser int) (models.User, error)
	GetUserByUserName(email string) (int, error)
	ParseToken(token string) (string, error)
	VerifyToken(accessToken string) (*token.Payload, error)
	GetMenuOptionAll(IdUser int) ([]models.RoleUser, error)
}

type RoleService interface {
	GetAll(query models.QueryParameter) ([]models.Role, int, error)
	GetById(idRol int) (models.Role, error)
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
