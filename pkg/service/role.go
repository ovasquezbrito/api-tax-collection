package service

import (
	
	"github.com/ovasquezbrito/tax-collection/pkg/models"
	"github.com/ovasquezbrito/tax-collection/pkg/repository"
)

type RolesService struct {
	repo repository.RoleRepository
}

func NewRoleService(repo repository.RoleRepository) *RolesService {
	return &RolesService{repo: repo}
}

func (s *RolesService) GetAll(query models.QueryParameter) ([]models.Role, int, error) {
	queryUpper := query.UpperCase()
	return s.repo.GetAll(*queryUpper)
}

func (s *RolesService) GetById(idRol int) (models.Role, error) {
	return s.repo.GetById(idRol)
}
