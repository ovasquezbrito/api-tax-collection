package service

import (
	baseapp "github.com/ovasquezbrito/base-app"
	"github.com/ovasquezbrito/base-app/pkg/repository"
)

type RolesService struct {
	repo repository.RoleRepository
}

func NewRoleService(repo repository.RoleRepository) *RolesService {
	return &RolesService{repo: repo}
}

func (s *RolesService) GetAll(query baseapp.QueryParameter) ([]baseapp.Role, int, error) {
	queryUpper := query.UpperCase()
	return s.repo.GetAll(*queryUpper)
}

func (s *RolesService) GetById(idRol int) (baseapp.Role, error) {
	return s.repo.GetById(idRol)
}
